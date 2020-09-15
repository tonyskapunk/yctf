package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func getFlag1(c *gin.Context) {
	f, err := getFlag(1)
	if err != nil {
		log.Fatal(err)
	}
	title := fmt.Sprintf("flag%v", f.ID)

	c.Header("x-yctf", f.Flag)
	render(c, gin.H{
		"title": title}, f.Template)
}
