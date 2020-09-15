package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func getFlag2(c *gin.Context) {
	f, err := getFlag(2)
	if err != nil {
		log.Fatal(err)
	}
	title := fmt.Sprintf("flag%v", f.ID)

	cookieName := fmt.Sprintf("flag%v", f.ID)
	c.SetCookie(cookieName, f.Flag, 300, "", "", false, false)
	render(c, gin.H{
		"title": title}, f.Template)
}
