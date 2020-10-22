package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getFlag5(c *gin.Context) {
	f, err := getFlag(5)
	if err != nil {
		log.Fatal(err)
	}
	title := fmt.Sprintf("flag%v", f.ID)

	auth, err := c.Cookie("auth")
	if err != nil || auth != "true" {
		c.SetCookie("auth", "false", 300, "", "", false, false)
		render(c, http.StatusUnauthorized, gin.H{"payload": f.Flag}, "401.html")
		return
	}

	if auth == "true" {
		render(c, http.StatusOK, gin.H{
			"payload": f.Flag,
			"title":   title}, f.Template)
		return
	}
}
