package main

import (
	"log"
	"os"
	"path/filepath"
	"torrsru/db"
	"torrsru/global"
	"torrsru/tgbot"
	"torrsru/web"

	"github.com/alexflint/go-arg"
)

func main() {
	var args struct {
		RebuildIndex bool   `default:"false" arg:"-r" help:"rebuild index and exit"`
		TMDBProxy    bool   `default:"false" arg:"--tmdb,env:TMDB_PROXY" help:"proxy for TMDB"`
		Port         string `default:"8094" arg:"-p,env:PORT" help:"port for http"`
		TGBotToken   string `default:"" arg:"--token,env:TG_TOKEN" help:"telegram bot token"`
		TGHost       string `default:"http://127.0.0.1:8081" arg:"--tgapi,env:TG_HOST" help:"telegram api host"`
		TSHost       string `default:"http://127.0.0.1:8090" arg:"--ts,env:TS_HOST" help:"TorrServer host"`
		DBHost       string `default:"http://62.112.8.193:9117" arg:"--db,env:DB_HOST" help:"External sync database host"`
		DBSync       int    `default:"20" arg:"--db-sync,env:DB_SYNC" help:"External database sync delay"`
		DBSyncRetry  int    `default:"10" arg:"--db-sync-retry,env:DB_SYNC_RETRY" help:"External database sync retry"`
	}
	arg.MustParse(&args)

	pwd := filepath.Dir(os.Args[0])
	pwd, _ = filepath.Abs(pwd)
	log.Println("PWD:", pwd)
	global.PWD = pwd

	global.TMDBProxy = args.TMDBProxy
	global.TSHost = args.TSHost

	global.DBHost = args.DBHost
	global.DBSync = args.DBSync
	global.DBSyncRetry = args.DBSyncRetry

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
