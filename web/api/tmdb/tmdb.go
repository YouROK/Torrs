package tmdb

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func TMDBAPI(c *gin.Context) {
	url := "https://api.themoviedb.org/" + strings.TrimPrefix(c.Request.RequestURI, "/tmdb/")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiI2MDBjYTlmZTIxYjJiMzk3MDY5OGRiNjJkNTc1OGUzOSIsInN1YiI6IjViMjJhYmJjMGUwYTI2NGRiODAxYzQ3MyIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ.NUqczMBJjZKycUKrRZxsNTrea8DS80GwX1T0VnY_O_A")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	defer res.Body.Close()

	c.Header("Cache-Control", "public, max-age=604800") //1 week
	c.Header("Etag", res.Header.Get("Etag"))
	c.DataFromReader(res.StatusCode, res.ContentLength, "application/javascript; charset=utf-8", res.Body, nil)
}

func TMDBIMG(c *gin.Context) {
	//url := "https://imagetmdb.com/" + strings.TrimPrefix(c.Request.RequestURI, "/tmdbimg/")
	url := "https://image.tmdb.org/" + strings.TrimPrefix(c.Request.RequestURI, "/tmdbimg/")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiI2MDBjYTlmZTIxYjJiMzk3MDY5OGRiNjJkNTc1OGUzOSIsInN1YiI6IjViMjJhYmJjMGUwYTI2NGRiODAxYzQ3MyIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ.NUqczMBJjZKycUKrRZxsNTrea8DS80GwX1T0VnY_O_A")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	defer res.Body.Close()

	c.Header("Cache-Control", "public, max-age=604800") //1 week
	c.Header("Etag", res.Header.Get("Etag"))
	c.DataFromReader(res.StatusCode, res.ContentLength, res.Header.Get("Content-Type"), res.Body, nil)
}
