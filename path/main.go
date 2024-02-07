package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// func main() {
// 	router := gin.Default()

// 	router.GET("/user/:name", getting)
// 	router.Run()
// }

//	func getting(c *gin.Context) {
//		name := c.Param("name")
//		c.String(http.StatusOK, "Hello %s", name)
//	}
// func main() {
// 	router := gin.Default()

// 	router.GET("/user/:name/*action", getting)
// 	router.Run()
// }

// func getting(c *gin.Context) {
// 	name := c.Param("name")
// 	action := c.Param("action")
// 	message := name + "is" + action
// 	c.String(http.StatusOK, message)
// }

func main() {
	router := gin.Default()
	router.GET("/user/:name/*action", getting)
	router.POST("/user/:name/*action", postting)
	router.GET("/user/groups", func(c *gin.Context) {
		c.String(http.StatusOK, "The available groups are [...]")
	})
	router.Run()
}

func postting(c *gin.Context) {

	b := c.FullPath() == "/user/:name/*action"
	c.String(http.StatusOK, "%t", b)
}
func getting(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	message := name + "is" + action
	c.String(http.StatusOK, message)
}
