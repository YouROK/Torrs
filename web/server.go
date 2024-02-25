package web

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"
	"log"
	"path/filepath"
	"torrsru/db/sync"
	"torrsru/web/api"
	"torrsru/web/gincert"
	"torrsru/web/global"
	"torrsru/web/static"
)

func Start() {
	go sync.StartSync()

	//gin.SetMode(gin.DebugMode)
	//gin.SetMode(gin.ReleaseMode)

	global.Route = gin.Default()
	global.Route.Use(gin.Recovery())
	static.RouteStaticFiles(global.Route)
	global.Route.LoadHTMLGlob(filepath.Join(global.PWD, "views/*.go.html"))
	api.SetRoutes(global.Route)

	if gin.Mode() == gin.DebugMode {
		global.Route.Run(":80")
	} else {
		am := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist("torrs.ru"),
			Cache:      autocert.DirCache(filepath.Join(global.PWD, "cache")),
		}

		err := gincert.RunWithManager(global.Route, &am)
		if err != nil {
			log.Println("Error start server:", err)
		}
	}

	global.Stoped = true
}
