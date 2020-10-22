package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func showIndex(c *gin.Context) {
	flags := getAllFlags()

	render(c, http.StatusOK, gin.H{
		"title":   "yCTF",
		"payload": flags}, "index.html")
}

// TODO:
func validateFlag(c *gin.Context) {
	render(c, http.StatusOK, gin.H{
		"payload": ""}, "validate.html")
}
