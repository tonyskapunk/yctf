package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getFlag7(c *gin.Context) {
	f, err := getFlag(7)
	if err != nil {
		log.Fatal(err)
	}
	title := fmt.Sprintf("flag%v", f.ID)

	type dataFlag struct {
		Flag   string `form:"flag"`
		Please string `form:"please"`
	}

	var d dataFlag
	isValid := false
	httpStatus := http.StatusBadRequest

	if c.BindQuery(&d) != nil {
		render(c, httpStatus, gin.H{
			"isValid": isValid,
			"title":   title}, f.Template)
		return
	}

	if d.Flag == title && d.Please != "" {
		isValid = true
		httpStatus = http.StatusOK
	}

	render(c, httpStatus, gin.H{
		"isValid": isValid,
		"flag":    f.Flag,
		"title":   title}, f.Template)
}
