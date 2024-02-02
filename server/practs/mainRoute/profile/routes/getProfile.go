package routes

import (
	"context"
	"net/http"

	"practs/mainRoute/pb"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Summary      Get Profile
// @Tags         profile
// @Description  Получение профиля юзера
// @Produce      json
// @Success      200   {object}  pb.GetProfileResponse
// @Failure      404   {object}  pb.GetProfileResponse
// @Router       /profile [get]
func GetProfile(ctx *gin.Context, c pb.ProfileServiceClient) {
	userId, _ := ctx.Get("userId")

	res, err := c.GetProfile(context.Background(), &pb.GetProfileRequest{
		UserId: userId.(int64),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
