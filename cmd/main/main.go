package main

import (
	"github.com/alexflint/go-arg"
	"log"
	"os"
	"path/filepath"
	"torrsru/db"
	"torrsru/global"
	"torrsru/tgbot"
	"torrsru/web"
)

func main() {
	var args struct {
		Port         string `default:"8094" arg:"-p" help:"port for http"`
		RebuildIndex bool   `default:"false" arg:"-r" help:"rebuild index and exit"`
		TMDBProxy    bool   `default:"false" arg:"--tmdb" help:"proxy for TMDB"`
		TGBotToken   string `default:"" arg:"--token" help:"telegram bot token"`
		TGHost       string `default:"http://127.0.0.1:8081" arg:"--tgapi" help:"telegram api host"`
		TSHost       string `default:"http://127.0.0.1:8090" arg:"--ts" help:"TorrServer host"`
	}
	arg.MustParse(&args)

	pwd := filepath.Dir(os.Args[0])
	pwd, _ = filepath.Abs(pwd)
	log.Println("PWD:", pwd)
	global.PWD = pwd

	global.TMDBProxy = args.TMDBProxy
	global.TSHost = args.TSHost

	db.Init()

	if args.RebuildIndex {
		err := db.RebuildIndex()
		if err != nil {
			log.Println("Rebuild index error:", err)
		} else {
			log.Println("Rebuild index success")
		}
		return
	}

	if args.TGBotToken != "" {
		if args.TGHost == "" {
			log.Println("Error telegram host is empty. Telegram api bot need for upload 2gb files")
			os.Exit(1)
		}
		err := tgbot.Start(args.TGBotToken, args.TGHost)
		if err != nil {
			log.Println("Start Telegram bot error:", err)
			os.Exit(1)
		}
	}
	web.Start(args.Port)
}
