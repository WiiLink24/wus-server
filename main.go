package main

import (
	"encoding/xml"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	config = &Config{}
)

func checkError(err error) {
	if err != nil {
		log.Fatalf("WUS Server has encountered an error! Reason: %v\n", err)
	}
}

func handleInquiry(c *gin.Context) {
	wiiNos := c.PostForm("chkno")

	wiiNoSplit := strings.Split(wiiNos, ",")

	c.Header("X-Wus-Host", config.WUSHost)
	c.Header("X-Result", "011")

	// Return a 1 for every Wii Number, this tells the game that all Wii Numbers are registered
	// TODO: Implement properly
	c.String(http.StatusOK, strings.Repeat("1", len(wiiNoSplit)))
}

func handleNotify(c *gin.Context) {
	c.Header("X-Wus-Host", config.WUSHost)
	c.Header("X-Result", "001")

	c.Status(http.StatusOK)
}

func main() {
	// Load the config
	rawConfig, err := os.ReadFile("./config.xml")
	checkError(err)

	err = xml.Unmarshal(rawConfig, config)
	checkError(err)

	if !config.DebugMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.POST("/inquiry", handleInquiry)
	r.POST("/notify", handleNotify)

	log.Fatal(r.Run(config.Address))
}
