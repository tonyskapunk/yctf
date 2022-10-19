package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func showIndex(c *gin.Context) {
	var score int

	flag_templates, err := filepath.Glob("./templates/flag*.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Only one session exists and is the default
	session := sessions.Default(c)

	s := session.Get("score")
	if s == nil {
		session.Set("score", score)
	}
	f := session.Get("total_flags")
	if f == nil {
		session.Set("total_flags", len(flag_templates))
	}
	if s == nil || f == nil {
		session.Save()
	}

	render(c, http.StatusOK, gin.H{
		"title":   "yCTF",
		"payload": flag_templates}, "index.html")
}

// TODO:
func validateFlag(c *gin.Context) {
	render(c, http.StatusOK, gin.H{
		"payload": ""}, "validate.html")
}
