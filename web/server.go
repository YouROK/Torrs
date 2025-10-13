package web

import (
	"embed"
	"html/template"
	"log"
	"strings"
	ss "sync"
	"time"
	"torrsru/db"
	"torrsru/global"
	"torrsru/web/api"
	"torrsru/web/static"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//go:embed views/*.go.html
var viewFS embed.FS

func Start(port string) {
	go db.StartSync()

	//gin.SetMode(gin.DebugMode)
	gin.SetMode(gin.ReleaseMode)

	corsCfg := cors.DefaultConfig()
	corsCfg.AllowAllOrigins = true
	corsCfg.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "X-Requested-With", "Accept", "Authorization"}

	global.Route = gin.New()
	global.Route.Use(gin.Recovery(), cors.New(corsCfg), blockUsers())
	static.RouteStaticFiles(global.Route)

	tmpl := template.Must(template.ParseFS(viewFS, "views/*.go.html"))
	global.Route.SetHTMLTemplate(tmpl)

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

		if strings.Contains(referer, "lamp") || strings.Contains(useragent, "lamp") {
			mu.Lock()
			c.Next()
			time.Sleep(time.Millisecond * 300)
			mu.Unlock()
			return
		}

		c.Next()
	}
}
