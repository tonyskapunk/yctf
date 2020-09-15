package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func getFlag0(c *gin.Context) {
	f, err := getFlag(0)
	if err != nil {
		log.Fatal(err)
	}
	title := fmt.Sprintf("flag%v", f.ID)

	render(c, gin.H{
		"payload": f.Flag,
		"title":   title}, f.Template)
}
