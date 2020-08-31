package main

import (
	"github.com/gin-gonic/gin"
)

func showIndex(c *gin.Context) {
	flags := getAllFlags()

	render(c, gin.H{
		"title":   "yCTF",
		"payload": flags}, "index.html")
}

// TODO:
func validateFlag(c *gin.Context) {
	render(c, gin.H{
		"payload": ""}, "validate.html")
}
