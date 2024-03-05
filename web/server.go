package web

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"path/filepath"
	"torrsru/db/sync"
	"torrsru/web/api"
	"torrsru/web/global"
	"torrsru/web/static"
)

func Start(port string) {
	go sync.StartSync()

	//gin.SetMode(gin.DebugMode)
	//gin.SetMode(gin.ReleaseMode)

	corsCfg := cors.DefaultConfig()
	corsCfg.AllowAllOrigins = true
	corsCfg.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "X-Requested-With", "Accept", "Authorization"}

	global.Route = gin.New()
	global.Route.Use(gin.Recovery(), cors.New(corsCfg))
	static.RouteStaticFiles(global.Route)
	global.Route.LoadHTMLGlob(filepath.Join(global.PWD, "views/*.go.html"))
	api.SetRoutes(global.Route)

	err := global.Route.Run(":" + port)
	if err != nil {
		log.Println("Error start server:", err)
	}

	//if gin.Mode() == gin.DebugMode {
	//	err := global.Route.Run(":80")
	//	if err != nil {
	//		log.Println("Error start server:", err)
	//	}
	//} else {
	//	am := autocert.Manager{
	//		Prompt:     autocert.AcceptTOS,
	//		HostPolicy: autocert.HostWhitelist("torrs.ru"),
	//		Cache:      autocert.DirCache(filepath.Join(global.PWD, "cache")),
	//	}
	//
	//	err := gincert.RunWithManager(global.Route, &am)
	//	if err != nil {
	//		log.Println("Error start server:", err)
	//	}
	//}

	global.Stoped = true
}
