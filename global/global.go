package global

import "github.com/gin-gonic/gin"

var (
	Route     *gin.Engine
	Stopped   = false
	PWD       = ""
	TMDBProxy = false
	TSHost    = ""

	DBHost      = ""
	DBSync      = 20
	DBSyncRetry = 10

	SendFromWeb func(initData, msg string) error
)
