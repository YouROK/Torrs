package pages

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func MainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "main.go.html", gin.Mode())
}
