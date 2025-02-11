package pages

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"torrsru/db"
	"strings"
)

func Search(c *gin.Context) {
	query := c.Query("query")
	if strings.TrimSpace(query) == "" {
		c.Status(http.StatusNoContent)
		return
	}

	trs, err := db.Search(query)
	if err != nil {
		log.Println("Error get from db list:", err)
		c.Status(http.StatusInternalServerError)
		return
	}

	buf, err := json.Marshal(trs)
	if err != nil {
		log.Println("Error marshal torr list:", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	if gin.Mode() == gin.ReleaseMode {
		estr := query + strconv.Itoa(len(trs))
		etag := fmt.Sprintf("%x", md5.Sum([]byte(estr)))
		c.Header("ETag", etag)
		c.Header("Cache-Control", "public, max-age=3600")
	}
	c.Data(200, "application/javascript; charset=utf-8", buf)
}
