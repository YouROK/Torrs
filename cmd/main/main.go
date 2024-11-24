package main

import (
	"github.com/alexflint/go-arg"
	"log"
	"os"
	"path/filepath"
	"torrsru/db"
	"torrsru/web"
	"torrsru/web/global"
)

func main() {
	var args struct {
		Port         string `default:"8094" arg:"-p" help:"port for http"`
		RebuildIndex bool   `default:"false" arg:"-r" help:"rebuild index and exit"`
	}
	arg.MustParse(&args)

	pwd := filepath.Dir(os.Args[0])
	pwd, _ = filepath.Abs(pwd)
	log.Println("PWD:", pwd)
	global.PWD = pwd

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

	web.Start(args.Port)
}
