package routing

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"ocserv/internal/providers/routes"
	"ocserv/pkg/config"
)

var router *gin.Engine

func Init() {
	router = gin.Default()
}

func RegisterRoutes() {
	configs := config.GetApp()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = configs.AllowOrigins
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")
	//corsConfig.AllowMethods = append(corsConfig.AllowMethods, "POST", "GET", "OPTIONS")
	router.Use(cors.New(corsConfig))
	apiGroup := router.Group("/api/v1/")
	routes.Register(apiGroup)
}

func Serve() {
	cfg := config.GetApp()
	RegisterRoutes()
	router.ForwardedByClientIP = true
	err := router.SetTrustedProxies(nil)
	if err != nil {
		log.Fatal(err)
	}
	server := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	fmt.Println("server running on " + server)
	err = router.Run(server)
	if err != nil {
		log.Fatal(err)
	}
}
