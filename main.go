package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	engine := gin.New()

	engine.GET("/", Test)

	engine.Run(":8081")
}

func Test(c *gin.Context) {
	c.String(http.StatusOK, "success")
}
