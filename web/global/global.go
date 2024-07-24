package global

import "github.com/gin-gonic/gin"

var (
	Route         *gin.Engine
	Stopped       = false
	PWD           = ""
	IsUpdateIndex = true
)
