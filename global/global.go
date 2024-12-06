package global

import "github.com/gin-gonic/gin"

var (
	Route     *gin.Engine
	Stopped   = false
	PWD       = ""
	TMDBProxy = false
	TSHost    = ""

	SendFromWeb func(initData, msg string) error
)
