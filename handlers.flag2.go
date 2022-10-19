package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getFlag2(c *gin.Context) {
	f, err := getFlag(2)
	if err != nil {
		log.Fatal(err)
	}
	title := fmt.Sprintf("flag%v", f.ID)

	cookieName := fmt.Sprintf("flag%v", f.ID)
	c.SetCookie(cookieName, f.Flag.String(), 300, "", "", false, false)
	render(c, http.StatusOK, gin.H{
		"title": title}, f.Template)
}
