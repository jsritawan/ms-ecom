package main

import (
	"flag"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/jsritawan/ms-ecom/internal/config"
)

var (
	cfg *config.Config
)

func init() {
	flag.Parse()

	cfg = config.LoadConfig(os.Getenv("CONFIG_PATH"))
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	port := make([]string, 0)
	if cfg.Server.Port != "" {
		port = append(port, ":"+cfg.Server.Port)
	}
	r.Run(port...)
}
