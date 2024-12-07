package torr

import (
	"fmt"
	tele "gopkg.in/telebot.v4"
	"math"
	"path/filepath"
	"strconv"
	"sync"
	"torrsru/db"
	"torrsru/tgbot/torr/state"
)

type DLQueue struct {
	id        int
	c         tele.Context
	hash      string
	fileID    string
	fileName  string
	updateMsg *tele.Message
}

var (
	queue   []*DLQueue
	mu, smu sync.Mutex
	isWork  bool
	idCount int
)

func Show(c tele.Context) error {
	msg := ""
	mu.Lock()
	for i, dlQueue := range queue {
		s := "#" + strconv.Itoa(i+1) + ":\n<b>Хэш:</b> <code>" + dlQueue.hash + "</code>\n<i>" + filepath.Base(dlQueue.fileName) + "</i>\n"
		if len(msg+s) > 1024 {
			err := c.Send(msg)
			if err != nil {
				return err
			}
			msg = ""
		}
		msg += s
	}
	mu.Unlock()
	if msg != "" {
		return c.Send("Очередь:\n" + msg)
	} else {
		return c.Send("Очередь пуста")
	}
}

func Add(c tele.Context, hash, fileID string) {
	mu.Lock()

	if len(queue) > 30 {
		c.Bot().Send(c.Recipient(), "Очередь переполнена, попробуйте попозже\n\nЭлементов в очереди:"+strconv.Itoa(len(queue)))
		mu.Unlock()
		return
	}

	idCount++
	if idCount > math.MaxInt {
		idCount = 0
	}

	dlQueue := &DLQueue{
		id:     idCount,
		c:      c,
		hash:   hash,
		fileID: fileID,
	}
	mu.Unlock()
	ti, _ := GetTorrentInfo(hash)
	if ti != nil {
		id, err := strconv.Atoi(dlQueue.fileID)
		if err == nil {
			file := ti.FindFile(id)
			if file != nil {
				idi := db.GetTGFileID(dlQueue.hash + "|" + dlQueue.fileID)
				if idi != "" {
					d := &tele.Document{}
					d.FileID = idi
					d.Caption = filepath.Base(file.Path)
					d.FileName = file.Path
					err = dlQueue.c.Send(d)
					if err == nil {
						return
					}
				}
				dlQueue.fileName = file.Path
			}
		}
	}
	mu.Lock()
	queue = append(queue, dlQueue)
	mu.Unlock()

	uMsg, _ := c.Bot().Send(c.Recipient(), "Подготовка к загрузке")

	dlQueue.updateMsg = uMsg
	go work()
	go sendStatus()
}

func Cancel(id int) {
	mu.Lock()
	defer mu.Unlock()
	for i, dlQueue := range queue {
		if dlQueue.id == id {
			dlQueue.c.Bot().Delete(dlQueue.updateMsg)
			queue = append(queue[:i], queue[i+1:]...)
			go sendStatus()
			return
		}
	}
}

func work() {
	smu.Lock()
	if isWork {
		smu.Unlock()
		return
	}
	isWork = true
	defer func() { isWork = false }()
	smu.Unlock()

	for {
		mu.Lock()
		if len(queue) == 0 {
			mu.Unlock()
			break
		}
		dlQueue := queue[0]
		queue = queue[1:]
		mu.Unlock()

		sendStatus()

		ti, _ := GetTorrentInfo(dlQueue.hash)
		var file *state.TorrentFileStat
		if ti != nil {
			id, _ := strconv.Atoi(dlQueue.fileID)
			file = ti.FindFile(id)
		}

		dlQueue.c.Bot().Notify(dlQueue.c.Recipient(), tele.UploadingVideo)

		caption := filepath.Base(file.Path)
		torrFile, err := NewTorrFile(dlQueue)
		if err != nil {
			dlQueue.c.Bot().Edit(dlQueue.updateMsg, err.Error())
			continue
		}

		d := &tele.Document{}
		d.File.FileReader = torrFile
		d.FileName = file.Path
		d.Caption = caption
		go func() {
			err := dlQueue.c.Send(d)
			torrFile.Close()
			if err != nil {
				fmt.Println("Ошибка загрузки в телеграм:", err)
				errstr := fmt.Sprintf("Ошибка загрузки в телеграм: %v", file.Path)
				dlQueue.c.Bot().Edit(dlQueue.updateMsg, errstr)
			} else {
				dlQueue.c.Bot().Delete(dlQueue.updateMsg)
				db.SaveTGFileID(dlQueue.hash+"|"+dlQueue.fileID, d.FileID)
			}
		}()
		<-torrFile.complete
	}
}

func sendStatus() {
	mu.Lock()
	defer mu.Unlock()
	for i, dlQueue := range queue {
		torrKbd := &tele.ReplyMarkup{}
		btnCancel := torrKbd.Data("Отмена", "downloadCancel", strconv.Itoa(dlQueue.id))
		rows := []tele.Row{torrKbd.Row(btnCancel)}
		torrKbd.Inline(rows...)

		msg := "Номер в очереди " + strconv.Itoa(i+1)
		if dlQueue.fileName != "" {
			msg += "\n<i>" + dlQueue.fileName + "</i>"
		}

		dlQueue.c.Bot().Edit(dlQueue.updateMsg, msg, torrKbd)
	}
}
