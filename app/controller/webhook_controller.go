package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"steam-notify/model/telegram"
)

func BadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, map[string]interface{}{"msg": fmt.Sprintf("Bad Request: %s", err.Error())})
	c.Abort()
}

func WebhookEntry(c *gin.Context) {
	bodyByte, err := io.ReadAll(c.Request.Body)
	if err != nil {
		BadRequest(c, err)
		return
	}
	var message telegram.BasicMessage
	err = json.Unmarshal(bodyByte, &message)
	if err != nil {
		BadRequest(c, err)
		return
	}
}
