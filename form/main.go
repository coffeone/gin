package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/form_post", postting)

	router.Run()
}

func postting(c *gin.Context) {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")

	c.JSON(http.StatusOK, gin.H{
		"status":  "posted",
		"message": message,
		"nick":    nick,
	})
}

// func main() {
// 	router := gin.Default()

// 	router.POST("/post", func(c *gin.Context) {

// 		id := c.Query("id")
// 		page := c.DefaultQuery("page", "0")
// 		name := c.PostForm("name")
// 		message := c.PostForm("message")

// 		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
// 	})
// 	router.Run(":8080")
// }
