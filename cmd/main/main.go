package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"torrsru/db"
	"torrsru/web"
	"torrsru/web/global"
)

func main() {
	pwd := filepath.Dir(os.Args[0])
	pwd, _ = filepath.Abs(pwd)
	log.Println("PWD:", pwd)
	global.PWD = pwd
	db.Init()

	port := flag.String("port", "8094", "port for web")

	if port == nil {
		p := "8094"
		port = &p
	}

	web.Start(*port)
}
