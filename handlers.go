package main

import (
	"net/http"

	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func showIndex(c *gin.Context) {
	var score int

	flags := challenges
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
		err := session.Save()
		// TODO: Identify proper path to take when a session save doesn't work
		log.Println("WARN unable to save session:", err)
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
