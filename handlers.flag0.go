package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func getFlag0(c *gin.Context) {
	f, err := getFlag(0)
	if err != nil {
		log.Fatal(err)
	}

	render(c, gin.H{
		"payload": f.Flag}, f.Template)
}
