package routes

import (
	"context"
	"net/http"

	"practs/mainRoute/pb"

	"github.com/gin-gonic/gin"
)

type LoginRequestBody struct {
	Logname  string `json:"logname"`
	Password string `json:"password"`
}

// @Summary      Authorization
// @Tags         auth
// @Description  Authorization in system
// @Produce      json
// @Param        main  body      pb.LoginRequest  true  "Login and pass"
// @Success      200   {object}  pb.LoginResponse
// @Failure      404   {object}  pb.LoginResponse
// @Router       /auth/login [post]
func Login(ctx *gin.Context, c pb.AuthServiceClient) {
	b := LoginRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Login(context.Background(), &pb.LoginRequest{
		Logname:  b.Logname,
		Password: b.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
