package routing

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	"log"
	_ "ocserv/docs"
	"ocserv/internal/providers/routes"
	"ocserv/pkg/config"
)

var router *gin.Engine

func Init() {
	router = gin.Default()
}

func GetRouter() *gin.Engine {
	return router
}

func RegisterRoutes() {
	configs := config.GetApp()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = configs.AllowOrigins
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")
	//corsConfig.AllowMethods = append(corsConfig.AllowMethods, "POST", "GET", "OPTIONS")
	r := GetRouter()
	r.Use(cors.New(corsConfig))
	apiGroup := r.Group("/api/v1/")
	routes.Register(apiGroup)
}

// Serve godoc
// @title           Ocserv Backend Example API
// @version         1.0
// @description     This is Ocserv Backend Api Doc server.
// @contact.name    API Support
// @host      localhost:8080
// @BasePath  /api/v1
func Serve() {
	cfg := config.GetApp()
	RegisterRoutes()
	router.ForwardedByClientIP = true
	err := router.SetTrustedProxies(nil)
	if err != nil {
		log.Fatal(err)
	}

	if cfg.Debug {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	server := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	fmt.Println("server running on " + server)
	err = router.Run(server)
	if err != nil {
		log.Fatal(err)
	}
}
