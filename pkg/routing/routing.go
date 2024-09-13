package routing

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"ocserv/internal/providers/routes"
	"ocserv/pkg/config"
)

var router *gin.Engine

func init() {
	router = gin.Default()
}

func Serve() {
	cfg := config.GetApp()
	routes.Register(router)
	server := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	fmt.Println("server running on " + server)
	err := router.Run(server)
	if err != nil {
		log.Fatal(err)
	}
}
