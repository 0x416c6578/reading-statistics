package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type somejson struct {
	SomeField      string `json:"someField"` // Ok this is cool
	SomeOtherField int    `json:"someOtherField"`
}

var someJsonExample somejson = somejson{SomeField: "hello", SomeOtherField: 123}

func main() {
	router := gin.Default()
	router.GET("somejson", getSomeJson)
	router.Run("localhost:8080")
}

func getSomeJson(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, someJsonExample)
}
