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
	if err != nil {
		c.SetCookie("auth", "false", 300, "", "", false, false)
		c.HTML(http.StatusUnauthorized, "401.html", gin.H{"payload": f.Flag})
		return
	}

	if auth == "true" {
		render(c, gin.H{
			"payload": f.Flag,
			"title":   title}, f.Template)
		return
	}
}
