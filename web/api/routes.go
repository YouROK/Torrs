package api

import (
	"github.com/gin-gonic/gin"
	"torrsru/global"
	"torrsru/web/api/pages"
	"torrsru/web/api/tgbot"
	"torrsru/web/api/tmdb"
)

func SetRoutes(r *gin.Engine) {
	r.GET("/", pages.MainPage)
	//r.GET("/robots.txt", pages.RobotsPage)
	r.GET("/search", pages.Search)
	if global.TMDBProxy {
		r.GET("/tmdb/*path", tmdb.TMDBAPI)
		r.GET("/tmdbimg/*path", tmdb.TMDBIMG)
	}
	r.POST("/sendbot", tgbot.SendBot)
}
