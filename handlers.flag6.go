package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getFlag6(c *gin.Context) {
	f, err := getFlag(6)
	if err != nil {
		log.Fatal(err)
	}
	title := fmt.Sprintf("flag%v", f.ID)

	validNetwork := "10.0.0.0/8"
	_, subnet, _ := net.ParseCIDR(validNetwork)

	var ip string
	isValid := false

	fwdIP := c.GetHeader("x-forwarded-for")
	srcIP := c.ClientIP()

	if fwdIP != "" {
		ip = fwdIP
	} else {
		ip = (srcIP)
	}

	if subnet.Contains(net.ParseIP(ip)) {
		isValid = true
	}

	render(c, http.StatusOK, gin.H{
		"isValid": isValid,
		"ip":      ip,
		"flag":    f.Flag,
		"title":   title}, f.Template)
}
