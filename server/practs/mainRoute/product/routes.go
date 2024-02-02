package product

import (
	"practs/mainRoute/auth"
	"practs/mainRoute/config"
	"practs/mainRoute/product/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/product")
	routes.Use(a.AuthRequired)
	routes.GET("/", svc.GetProduct)
}

func (svc *ServiceClient) GetProduct(ctx *gin.Context) {
	routes.GetProduct(ctx, svc.Client)
}
