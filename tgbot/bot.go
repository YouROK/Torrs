package tgbot

import (
	"errors"
	"fmt"
	initdata "github.com/telegram-mini-apps/init-data-golang"
	tele "gopkg.in/telebot.v4"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"torrsru/global"
	"torrsru/tgbot/torr"
)

func Start(token, host string) error {
	pref := tele.Settings{
		URL:       host,
		Token:     token,
		Poller:    &tele.LongPoller{Timeout: 5 * time.Minute},
		ParseMode: tele.ModeHTML,
		Client:    &http.Client{Timeout: 5 * time.Minute},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		return err
	}

	b.Handle("help", help)
	b.Handle("Help", help)
	b.Handle("/help", help)
	b.Handle("/Help", help)
	b.Handle("/start", help)

	b.Handle("/queue", torr.ShowQueue)

	b.Handle("/id", func(c tele.Context) error {
		return c.Send(fmt.Sprintf("%v %v %v %v", c.Sender().ID, c.Sender().Username, c.Sender().FirstName, c.Sender().LastName))
	})
	b.Handle("/exit", func(c tele.Context) error {
		if c.Sender().ID == 140045144 {
			c.Send("Exit")
			os.Exit(0)
		}
		return nil
	})

	b.Handle(tele.OnText, func(c tele.Context) error {
		txt := c.Text()
		if strings.HasPrefix(strings.ToLower(txt), "magnet:") || isHash(txt) {
			return infoTorrent(c, c.Text())
		} else if c.Message().ReplyTo != nil && c.Message().ReplyTo.ReplyMarkup != nil && len(c.Message().ReplyTo.ReplyMarkup.InlineKeyboard) > 0 {
			data := c.Message().ReplyTo.ReplyMarkup.InlineKeyboard[0][0].Data
			if strings.HasPrefix(strings.ToLower(data), "\fall|") {
				hash := strings.TrimPrefix(data, "\fall|")
				from, to, err := ParseRange(c.Message().Text)
				if err != nil {
					c.Send("Ошибка, нужно указывать числа, пример: 2-12")
					return err
				}
				torr.AddRange(c, hash, from, to)
			}
			return nil
		} else {
			return c.Send("Вставьте магнет/хэш торрента или нажмите на поиск\n\nВ окне поиска введите название и в списке торрентов нажмите на +\n\nУчтите что файл не должен превышать 2гб это лимит телеграмма на отправку файлов")
		}
	})

	b.Handle(tele.OnCallback, func(c tele.Context) error {
		args := c.Args()
		if len(args) > 0 {
			if args[0] == "\ffile" || args[0] == "\fall" {
				return getTorrent(c)
			}
			if args[0] == "\ftorr" {
				return infoTorrent(c, args[1])
			}
			if args[0] == "\fcancel" {
				if num, err := strconv.Atoi(args[1]); err == nil {
					torr.Cancel(num)
					c.Bot().Delete(c.Callback().Message)
					return nil
				}
			}
		}
		return errors.New("Ошибка кнопка не распознана")
	})

	global.SendFromWeb = func(initDataUser, msg string) error {
		err := initdata.Validate(initDataUser, token, time.Duration(0))
		if err != nil {
			return errors.New("Error auth user")
		}
		data, err := initdata.Parse(initDataUser)
		if err != nil {
			return errors.New("Error parse user data")
		}
		chat, err := b.ChatByID(data.User.ID)
		if err != nil {
			return errors.New("Chat with user not found")
		}
		u := tele.Update{
			Message: &tele.Message{
				Sender: &tele.User{
					ID:           data.User.ID,
					FirstName:    data.User.FirstName,
					LastName:     data.User.LastName,
					Username:     data.User.Username,
					LanguageCode: data.User.LanguageCode,
					IsBot:        data.User.IsBot,
					IsPremium:    data.User.IsPremium,
					AddedToMenu:  data.User.AddedToAttachmentMenu,
				},
				Unixtime: time.Now().Unix(),
				Chat:     chat,
				Text:     msg,
			},
		}
		c := b.NewContext(u)
		return infoTorrent(c, msg)
	}

	torr.Start()
	go b.Start()

	return nil
}

func help(c tele.Context) error {
	return c.Send("Для поиска нажмите кнопку \"Поиск\", в списке нажать <b>+</b> для добавления на скачивание\n" +
		"Так же можно вставить магнет или хэш торрента\n" +
		"Лимит телеграма на загружаемый файл 2гб, выбирайте торренты, где файл будет меньше 2гб")
}

func ParseRange(rng string) (int, int, error) {
	parts := strings.Split(rng, "-")

	if len(parts) != 2 {
		return -1, -1, errors.New("Неверный формат строки")
	}

	num1, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err1 != nil {
		return -1, -1, err1
	}

	num2, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err2 != nil {
		return -1, -1, err2
	}
	return num1, num2, nil
}
