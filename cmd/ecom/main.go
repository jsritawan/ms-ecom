package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/jsritawan/ms-ecom/internal/config"
	"github.com/jsritawan/ms-ecom/internal/handler"
	"github.com/jsritawan/ms-ecom/internal/httpadapter"
	"github.com/jsritawan/ms-ecom/internal/storage"
)

var (
	cfg        *config.Config
	connection *storage.Storage
	service    *handler.Service
	adapter    *httpadapter.Adapter
)

func init() {
	flag.Parse()

	cfg = config.LoadConfig(os.Getenv("CONFIG_PATH"))
	connection = storage.New(&cfg.DB)
	connection.AutoMigrate()
	service = handler.New(connection)
	adapter = httpadapter.New(service)
}

func main() {
	g := gin.Default()

	api := g.Group("/api")

	v1 := api.Group("/v1")
	v1.GET("/health-check", func(c *gin.Context) {
		err := connection.HeathCheck()
		if err != nil {
			c.AbortWithStatus(http.StatusServiceUnavailable)
			return
		}

		c.Status(http.StatusNoContent)
	})

	v1.POST("/sign-up", adapter.AddUser)

	port := make([]string, 0)
	if cfg.Server.Port != "" {
		port = append(port, ":"+cfg.Server.Port)
	}
	g.Run(port...)
}
