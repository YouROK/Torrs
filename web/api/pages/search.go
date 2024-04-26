package pages

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
	"torrsru/db/search"
)

func Search(c *gin.Context) {
	query := c.Query("query")
	if strings.TrimSpace(query) == "" {
		c.Status(http.StatusNoContent)
		return
	}
	_, accurate := c.GetQuery("accurate")
	trs := search.Find(query, accurate)
	c.Header("Cache-Control", "public, max-age=3600")
	buf, err := json.Marshal(trs)
	if err != nil {
		log.Println("Error marshal torr list:", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	estr := query + strconv.Itoa(len(trs))
	etag := fmt.Sprintf("%x", md5.Sum([]byte(estr)))
	c.Header("ETag", etag)
	c.Data(200, "application/javascript; charset=utf-8", buf)
}
