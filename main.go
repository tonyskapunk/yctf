package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/tonyskapunk/yctf/pkg/flag"
)

const (
	defaultAuthKey         = "Jsj6XNRTNXv66VgkAEFGhZXbFH3jMfE9m36VXFEYv5"
	defaultSessionName     = "yctf-session"
	defaultFlagFactorySeed = "abcdefghijklmnopqrstuvwxyz"
	defaultFlagFile        = "flags.json"
)

var router *gin.Engine
var challenges []challenge

func main() {
	router := setupRouter()
	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	setupFlagFactory()
	generateFlags()
}
func setupFlagFactory() {
	// TODO: randomize this, or accept the seed from config.
	flag.SetFactory(flag.NewFactory(defaultFlagFactorySeed))
}

func generateFlags() {
	var ch []challenge

	data, err := os.ReadFile(defaultFlagFile)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal([]byte(data), &ch)
	if err != nil {
		log.Fatal(err)
	}

	for i, c := range ch {
		// For now, use the template as the input because it's required
		// for all challenges.
		ch[i].Flag = flag.NewFlag([]byte(c.Template))
	}

	challenges = ch
}

func setupRouter() *gin.Engine {
	authKey := os.Getenv("YCTF_AUTHKEY")
	if authKey == "" {
		authKey = defaultAuthKey
	}
	sessionName := os.Getenv("YCTF_SESSION_NAME")
	if sessionName == "" {
		sessionName = defaultSessionName
	}
	store := cookie.NewStore([]byte(authKey))
	store.Options(sessions.Options{SameSite: http.SameSiteStrictMode})

	router = gin.Default()
	router.Use(sessions.Sessions(sessionName, store))
	router.LoadHTMLGlob("templates/*")
	initializeRoutes()
	return router
}

func render(c *gin.Context, httpStatus int, data gin.H, templateName string) {
	reqHeader := c.Request.Header.Get("Content-Type")
	if reqHeader == "" {
		reqHeader = c.Request.Header.Get("Accept")
	}

	switch reqHeader {
	case "application/json":
		// Respond with JSON
		c.JSON(httpStatus, data["payload"])
	case "application/x-www-form-urlencoded":
		// Respond with JSON
		c.JSON(httpStatus, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(httpStatus, data["payload"])
	default:
		// Respond with HTML
		session := sessions.Default(c)
		data["score"] = session.Get("score")
		data["flags"] = session.Get("flags")
		c.HTML(httpStatus, templateName, data)
	}

}
