package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getFlag1(c *gin.Context) {
	f, err := getFlag(1)
	if err != nil {
		log.Fatal(err)
	}
	title := fmt.Sprintf("flag%v", f.ID)

	c.Header("x-yctf", f.Flag.String())
	render(c, http.StatusOK, gin.H{
		"title": title}, f.Template)
}
