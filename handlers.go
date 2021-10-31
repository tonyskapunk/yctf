package main

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func showIndex(c *gin.Context) {
	var score int

	flags := getAllFlags()
	// Only one session exists and is the default
	session := sessions.Default(c)

	s := session.Get("score")
	if s == nil {
		session.Set("score", score)
	}
	f := session.Get("flags")
	if f == nil {
		session.Set("flags", len(flags))
	}
	if s == nil || f == nil {
		session.Save()
	}

	render(c, http.StatusOK, gin.H{
		"title":   "yCTF",
		"payload": flags}, "index.html")
}

// TODO:
func validateFlag(c *gin.Context) {
	render(c, http.StatusOK, gin.H{
		"payload": ""}, "validate.html")
}
