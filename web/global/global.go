package global

import "github.com/gin-gonic/gin"

var (
	Route  *gin.Engine
	Stoped bool   = false
	PWD    string = ""
)
