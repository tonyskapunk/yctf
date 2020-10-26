package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router := setupRouter()
	router.Run()
}

func setupRouter() *gin.Engine {
	router = gin.Default()
	router.LoadHTMLGlob("templates/*")
	initializeRoutes()
	return router
}

func render(c *gin.Context, httpStatus int, data gin.H, templateName string) {
	reqHeader := c.Request.Header.Get("Content-Type")
	if reqHeader == "" {
		reqHeader = c.Request.Header.Get("Accept")
	}

	switch reqHeader {
	case "application/json":
		// Respond with JSON
		c.JSON(httpStatus, data["payload"])
	case "application/x-www-form-urlencoded":
		// Respond with JSON
		c.JSON(httpStatus, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(httpStatus, data["payload"])
	default:
		// Respond with HTML
		c.HTML(httpStatus, templateName, data)
	}

}
