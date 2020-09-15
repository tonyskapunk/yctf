package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

func getFlag4(c *gin.Context) {
	f, err := getFlag(4)
	if err != nil {
		log.Fatal(err)
	}
	title := fmt.Sprintf("flag%v", f.ID)

	ua := c.GetHeader("User-Agent")
	matched, err := regexp.Match(`.*bot.*`, []byte(ua))
	if err != nil {
		log.Fatal(err)
	}

	c.HTML(http.StatusOK, f.Template, gin.H{
		"isBot": matched,
		"ua":    ua,
		"flag":  f.Flag,
		"title": title})
}
