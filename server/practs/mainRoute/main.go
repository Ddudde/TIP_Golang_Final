package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"practs/mainRoute/auth"
	"practs/mainRoute/config"
	"practs/mainRoute/product"
	"practs/mainRoute/profile"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	_ "practs/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	gen "practs/mainRoute/pb"
)

// @title           MainRouteService
// @version         1.0
// @description     Route in HTTP to Services(AuthM, ProductM, ProfileM).\n\nMainRoute: doc.json\nAuthService: http://localhost:3000/docs1/mainRoute/pb/auth.swagger.json\nProductService: http://localhost:3000/docs1/mainRoute/pb/product.swagger.json\nProfileService: http://localhost:3000/docs1/mainRoute/pb/profile.swagger.json
// @schemes      HTTP
// @host      localhost:3000
// @BasePath  /

// @securityDefinitions.apikey  ApiKeyAuth
// @in header
// @name Authorization

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	newServiceToSwag(r, "50051", []string{"/authM"})
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Static("/docs1", "./docs")

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	profile.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}

func newServiceToSwag(router *gin.Engine, port string, urls []string) http.Handler {
	opts := []grpc.DialOption{grpc.WithInsecure()}

	gwmux := runtime.NewServeMux()
	if err := gen.RegisterAuthServiceHandlerFromEndpoint(context.Background(), gwmux, "localhost:"+port, opts); err != nil {
		log.Fatalln("Failed at swag"+port, err)
	}

	for _, url := range urls {
		router.Group(url+"/*{grpc_gateway}").Any("", gin.WrapH(gwmux))
	}

	return gwmux
}
