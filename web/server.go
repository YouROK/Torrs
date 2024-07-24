package web

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"path/filepath"
	"strings"
	ss "sync"
	"time"
	"torrsru/db/sync"
	"torrsru/web/api"
	"torrsru/web/global"
	"torrsru/web/static"
)

func Start(port string) {
	go sync.StartSync()

	//gin.SetMode(gin.DebugMode)
	gin.SetMode(gin.ReleaseMode)

	corsCfg := cors.DefaultConfig()
	corsCfg.AllowAllOrigins = true
	corsCfg.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "X-Requested-With", "Accept", "Authorization"}

	global.Route = gin.New()
	global.Route.Use(gin.Recovery(), cors.New(corsCfg), blockUsers())
	static.RouteStaticFiles(global.Route)
	global.Route.LoadHTMLGlob(filepath.Join(global.PWD, "views/*.go.html"))
	api.SetRoutes(global.Route)

	err := global.Route.Run(":" + port)
	if err != nil {
		log.Println("Error start server:", err)
	}

	global.Stopped = true
}

func blockUsers() gin.HandlerFunc {
	var mu ss.Mutex
	return func(c *gin.Context) {
		referer := strings.ToLower(c.Request.Referer())
		useragent := strings.ToLower(c.Request.UserAgent())

		if strings.Contains(referer, "lampishe") || strings.Contains(useragent, "lampishe") || strings.Contains(referer, "lampa") || strings.Contains(useragent, "lampa") {
			mu.Lock()
			c.Next()
			time.Sleep(time.Millisecond * 300)
			mu.Unlock()
			return
		}

		c.Next()
	}
}
