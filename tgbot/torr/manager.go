package torr

import (
	"errors"
	"fmt"
	tele "gopkg.in/telebot.v4"
	"math"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
	"torrsru/db"
	"torrsru/tgbot/torr/state"
)

type Worker struct {
	id          int
	c           tele.Context
	msg         *tele.Message
	torrentHash string
	isCancelled bool
	cmd         string
	ti          *state.TorrentStatus
}

type Manager struct {
	queue     []*Worker
	working   map[int]*Worker
	ids       int
	wrkSync   sync.Mutex
	queueLock sync.Mutex
}

func (m *Manager) Start() {
	m.working = make(map[int]*Worker)
	go m.work()
}

func (m *Manager) Add(c tele.Context, hash, cmd string) {
	m.queueLock.Lock()
	defer m.queueLock.Unlock()

	if len(m.queue) > 50 {
		c.Bot().Send(c.Recipient(), "Очередь переполнена, попробуйте попозже\n\nЭлементов в очереди:"+strconv.Itoa(len(m.queue)))
		return
	}

	m.ids++
	if m.ids > math.MaxInt {
		m.ids = 0
	}

	msg, _ := c.Bot().Send(c.Recipient(), "<b>Подключение к торренту</b>\n<code>"+hash+"</code>")
	ti, _ := GetTorrentInfo(hash)
	if ti == nil {
		c.Bot().Edit(msg, "Ошибка при подключении к торренту <code>"+hash+"</code>")
		return
	}

	w := &Worker{
		id:          m.ids,
		c:           c,
		torrentHash: hash,
		cmd:         cmd,
		msg:         msg,
		ti:          ti,
	}

	m.queue = append(m.queue, w)
}

func (m *Manager) Cancel(id int) {
	m.queueLock.Lock()
	defer m.queueLock.Unlock()
	var rem []int
	for i, w := range m.queue {
		if w.id == id {
			w.isCancelled = true
			w.c.Bot().Delete(w.msg)
			rem = append(rem, i)
			return
		}
	}
	for _, i := range rem {
		m.queue = append(m.queue[:i], m.queue[i+1:]...)
	}
	if wrk, ok := m.working[id]; ok {
		wrk.isCancelled = true
		return
	}
}

func (m *Manager) work() {
	for {
		m.queueLock.Lock()
		if len(m.working) > 3 {
			m.queueLock.Unlock()
			m.sendQueueStatus()
			time.Sleep(time.Second)
			continue
		}
		if len(m.queue) == 0 {
			m.queueLock.Unlock()
			time.Sleep(time.Second)
			continue
		}
		wrk := m.queue[0]
		m.queue = m.queue[1:]
		m.working[wrk.id] = wrk
		m.queueLock.Unlock()

		m.sendQueueStatus()

		go func() {
			if wrk.cmd == "all" {
				loadingAll(wrk)
			} else {
				loading(wrk)
			}
			m.queueLock.Lock()
			delete(m.working, wrk.id)
			m.queueLock.Unlock()
		}()
	}
}

func (m *Manager) sendQueueStatus() {
	m.queueLock.Lock()
	defer m.queueLock.Unlock()
	for i, wrk := range m.queue {
		if wrk.msg == nil {
			continue
		}
		torrKbd := &tele.ReplyMarkup{}
		torrKbd.Inline([]tele.Row{torrKbd.Row(torrKbd.Data("Отмена", "cancel", strconv.Itoa(wrk.id)))}...)

		msg := "Номер в очереди " + strconv.Itoa(i+1)
		if wrk.cmd == "all" {
			msg += "\n<i>Скачать все файлы</i>"
		} else if strings.HasPrefix(wrk.cmd, "file:") {
			id, _ := strconv.Atoi(strings.TrimPrefix(wrk.cmd, "file:"))
			file := wrk.ti.FindFile(id)
			msg += "\n<i>" + file.Path + "</i>"
		}

		wrk.c.Bot().Edit(wrk.msg, msg, torrKbd)
	}
}

func loadingAll(wrk *Worker) {
	iserr := false
	for i, file := range wrk.ti.FileStats {
		if wrk.isCancelled {
			return
		}

		caption := filepath.Base(file.Path)
		torrFile, err := NewTorrFile(wrk, file)
		if err != nil {
			wrk.c.Bot().Edit(wrk.msg, err.Error())
			return
		}

		var wa sync.WaitGroup
		wa.Add(1)
		complete := false
		go func() {
			for !complete {
				updateLoadStatus(wrk, torrFile, i+1, len(wrk.ti.FileStats))
				time.Sleep(1 * time.Second)
			}
			wa.Done()
		}()

		tgfid := db.GetTGFileID(wrk.torrentHash + "|" + strconv.Itoa(file.Id))
		d := &tele.Document{}
		d.FileName = file.Path
		d.Caption = caption
		if tgfid != "" {
			d.FileID = tgfid
		} else {
			d.File.FileReader = torrFile
		}

		err = wrk.c.Send(d)
		complete = true
		wa.Wait()
		torrFile.Close()
		if errors.Is(err, ERR_STOPPED) {
			wrk.c.Bot().Delete(wrk.msg)
		} else if err != nil {
			fmt.Println("Ошибка загрузки в телеграм:", err)
			errstr := fmt.Sprintf("Ошибка загрузки в телеграм: %v", file.Path)
			wrk.c.Bot().Edit(wrk.msg, errstr)
			iserr = true
			return
		} else {
			db.SaveTGFileID(wrk.torrentHash+"|"+strconv.Itoa(file.Id), d.FileID)
		}
	}
	if !iserr {
		wrk.c.Bot().Delete(wrk.msg)
	}
}

func loading(wrk *Worker) {
	var file *state.TorrentFileStat
	id, _ := strconv.Atoi(strings.TrimPrefix(wrk.cmd, "file:"))
	file = wrk.ti.FindFile(id)

	caption := filepath.Base(file.Path)
	torrFile, err := NewTorrFile(wrk, file)
	if err != nil {
		wrk.c.Bot().Edit(wrk.msg, err.Error())
		return
	}

	var wa sync.WaitGroup
	wa.Add(1)
	complete := false
	go func() {
		for !complete {
			updateLoadStatus(wrk, torrFile, 0, 0)
			time.Sleep(1 * time.Second)
		}
		wa.Done()
	}()

	d := &tele.Document{}
	d.File.FileReader = torrFile
	d.FileName = file.Path
	d.Caption = caption

	err = wrk.c.Send(d)
	complete = true
	wa.Wait()
	torrFile.Close()
	if errors.Is(err, ERR_STOPPED) {
		wrk.c.Bot().Delete(wrk.msg)
	} else if err != nil {
		fmt.Println("Ошибка загрузки в телеграм:", err)
		errstr := fmt.Sprintf("Ошибка загрузки в телеграм: %v", file.Path)
		wrk.c.Bot().Edit(wrk.msg, errstr)
	} else {
		wrk.c.Bot().Delete(wrk.msg)
		db.SaveTGFileID(wrk.torrentHash+"|"+strconv.Itoa(file.Id), d.FileID)
	}
}
