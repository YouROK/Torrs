package web

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path/filepath"
	"strings"
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

	global.Stoped = true
}

func blockUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		referer := strings.ToLower(c.Request.Referer())
		useragent := strings.ToLower(c.Request.UserAgent())

		if strings.Contains(referer, "lampishe") || strings.Contains(useragent, "lampishe") {
			if strings.Contains(c.Request.RequestURI, "/tmdbimg/") {
				c.Redirect(http.StatusMovedPermanently, "http://releases.yourok.ru/torr/fake.png")
				return
			}

			if strings.Contains(c.Request.RequestURI, "/tmdb/") {
				c.Request.RequestURI = strings.ReplaceAll(c.Request.RequestURI, "language=ru", "language=bg")
				c.Request.RequestURI = strings.ReplaceAll(c.Request.RequestURI, "language=ua", "language=bg")
			}
		}

		c.Next()
	}
}
