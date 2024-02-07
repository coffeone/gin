package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var secrets = gin.H{
	"aaa": gin.H{"email": "aaa@mail.com", "phone": "123456"},
	"bbb": gin.H{"email": "bbb@mail.com", "phone": "888888"},
	"ccc": gin.H{"email": "ccc@mail.com", "phone": "333333"},
}

func main() {
	r := gin.Default()
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"aaa": "bar",
		"bbb": "hello",
		"ccc": "edssd",
		"ddd": "5262",
	}))

	authorized.GET("/secrets", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})
	r.Run(":8080")
}
