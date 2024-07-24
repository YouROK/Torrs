package pages

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"torrsru/db/search"
	"torrsru/models/fdb"
)

func Search(c *gin.Context) {
	query := c.Query("query")
	_, accurate := c.GetQuery("accurate")
	_, byword := c.GetQuery("byword")
	var trs []*fdb.Torrent
	if byword {
		trs = search.FindTitle(query)
	} else {
		trs = search.FindName(query, accurate)
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
