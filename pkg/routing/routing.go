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

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func Serve() {
	cfg := config.GetApp()
	RegisterRoutes()
	router.ForwardedByClientIP = true
	err := router.SetTrustedProxies(nil)
	if err != nil {
		log.Fatal(err)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	server := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	fmt.Println("server running on " + server)
	err = router.Run(server)
	if err != nil {
		log.Fatal(err)
	}
}
