package pages

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var modTime time.Time

func init() {
	modTime = time.Now()
}

func RobotsPage(c *gin.Context) {
	content := `
User-agent: *
Disallow:
`
	if c.Request.Method != "GET" && c.Request.Method != "HEAD" {
		status := http.StatusOK
		if c.Request.Method != "OPTIONS" {
			status = http.StatusMethodNotAllowed
		}
		c.Header("Allow", "GET,HEAD,OPTIONS")
		c.AbortWithStatus(status)
		return
	}
	c.Header("Content-Type", "text/plain")
	http.ServeContent(c.Writer, c.Request, "robots.txt", modTime, bytes.NewReader([]byte(content)))
}
