package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"webhook-service/webhook"
)

func setRouter(router *gin.Engine) {
	router.GET("/", getIndex)
	router.POST("/webhook/*path", postWebhook)
}

func getIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func postWebhook(c *gin.Context) {
	url := c.Param("path")
	url = strings.TrimPrefix(url, "/")
	w, found := webhook.GetByURL(url)
	msg := "Not found."
	if found {
		msg = "ok"
		w.Execute()
	}
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
	})
}
