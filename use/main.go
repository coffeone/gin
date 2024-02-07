package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// router.GET("/someGet", getting)
	router.POST("/somePost", postting)
	router.Run()
}

//	func getting(c *gin.Context) {
//		c.JSON(
//			http.StatusOK,
//			gin.H{
//				"qaq": "welcome server 01",
//			},
//		)
//	}
func postting(c *gin.Context) {
	username := c.PostForm("username")
	c.JSON(
		http.StatusOK,
		gin.H{"username": username},
	)
}
