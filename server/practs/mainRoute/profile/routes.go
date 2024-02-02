package profile

import (
	"practs/mainRoute/auth"
	"practs/mainRoute/config"
	"practs/mainRoute/profile/routes"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/profile")
	routes.Use(a.AuthRequired)
	routes.GET("/", svc.GetProfile)
}

func (svc *ServiceClient) GetProfile(ctx *gin.Context) {
	routes.GetProfile(ctx, svc.Client)
}
