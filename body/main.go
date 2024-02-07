package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type formA struct {
	Foo string `json:"foo" xml: "foo" binding: "required"`
}

type formB struct {
	Bar string `json:"bar" xml: "bar" binding: "required"`
}

func SomeHandle(c *gin.Context) {
	objA := formA{}
	objB := formB{}

	if errA := c.ShouldBind(&objA); errA == nil {
		c.String(http.StatusOK, `the body should be formA`)
	} else if errB := c.ShouldBind(&objB); errB == nil {
		c.String(http.StatusOK, `the body should be formB`)
	} else {
		return
	}
}
func main() {
	SomeHandle()
}
