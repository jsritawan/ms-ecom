package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/jsritawan/ms-ecom/internal/config"
	"github.com/jsritawan/ms-ecom/internal/storage"
)

var (
	cfg    *config.Config
	dbconn *storage.Storage
)

func init() {
	flag.Parse()

	cfg = config.LoadConfig(os.Getenv("CONFIG_PATH"))
	dbconn = storage.New(&cfg.DB)
}

func main() {
	g := gin.Default()

	api := g.Group("/api")

	v1 := api.Group("/v1")
	v1.GET("/health-check", func(c *gin.Context) {
		err := dbconn.HeathCheck()
		if err != nil {
			c.AbortWithStatus(http.StatusServiceUnavailable)
			return
		}

		c.Status(http.StatusNoContent)
	})

	port := make([]string, 0)
	if cfg.Server.Port != "" {
		port = append(port, ":"+cfg.Server.Port)
	}
	g.Run(port...)
}
