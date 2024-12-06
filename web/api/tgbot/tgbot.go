package tgbot

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"torrsru/global"
)

type TGSendData struct {
	InitData string `json:"init_data"`
	Magnet   string `json:"magnet"`
}

func SendBot(c *gin.Context) {
	req := TGSendData{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if global.SendFromWeb != nil {
		err = global.SendFromWeb(req.InitData, req.Magnet)
		if err == nil {
			c.Status(200)
		} else {
			c.AbortWithError(http.StatusBadRequest, err)
		}
	} else {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}
