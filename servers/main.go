package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group //タスクのどれかが失敗した場合は処理全体を失敗したものとして扱いたい
)

func handler(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"code":  http.StatusOK,
			"error": "welcome server 01",
		},
	)
}
func router01() http.Handler {

	e := gin.New()
	e.Use(gin.Recovery())

	e.GET("/", handler)
	return e
}
