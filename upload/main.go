package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.MaxMultipartMemory = 8 << 20
	router.POST("/upload", postting)
	router.Run()
}

func postting(c *gin.Context) {
	file, _ := c.FormFile("file")
	log.SetPrefix(file.Filename)
	dst := "uploads/" + file.Filename
	c.SaveUploadedFile(file, dst)
	c.String(http.StatusOK, fmt.Sprintf("''%s uploaded!", file.Filename))
}
