package tgbot

import (
	"errors"
	"github.com/dustin/go-humanize"
	tele "gopkg.in/telebot.v4"
	"path/filepath"
	"strconv"
	"strings"
	"torrsru/tgbot/torr"
)

func infoTorrent(c tele.Context, magnet string) error {
	msg, _ := c.Bot().Send(c.Recipient(), "Подключение к торренту: <code>"+magnet+"</code>")
	ti, err := torr.GetTorrentInfo(magnet)
	if err != nil {
		_, err = c.Bot().Edit(msg, "Ошибка при подключении к торренту <code>"+magnet+"</code>")
		return err
	}

	c.Bot().Delete(msg)

	txt := "<b>" + ti.Title + "</b>\n" +
		"<code>" + ti.Hash + "</code>"

	filesKbd := &tele.ReplyMarkup{}
	var files []tele.Row

	i := len(txt)
	for _, f := range ti.FileStats {
		btn := filesKbd.Data(filepath.Base(f.Path)+" "+humanize.Bytes(uint64(f.Length)), "file", ti.Hash, strconv.Itoa(f.Id))
		files = append(files, filesKbd.Row(btn))
		if i+len(txt) > 1024 {
			filesKbd := &tele.ReplyMarkup{}
			filesKbd.Inline(files...)
			c.Send(txt, filesKbd)
			files = files[:0]
			i = len(txt)
		}
		i += len(filepath.Base(f.Path) + " " + humanize.Bytes(uint64(f.Length)))
	}
	if len(files) > 0 {
		filesKbd.Inline(files...)
		return c.Send(txt, filesKbd)
	}
	return nil
}

func getTorrent(c tele.Context) error {
	args := c.Args()
	if len(args) != 3 {
		return errors.New("Ошибка не верные данные")
	}

	torr.Add(c, args[1], args[2])

	return nil
}

func isHash(txt string) bool {
	if len(txt) == 40 {
		for _, c := range strings.ToLower(txt) {
			switch c {
			case 'a', 'b', 'c', 'd', 'e', 'f', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			default:
				return false
			}
		}
		return true
	}
	return false
}
