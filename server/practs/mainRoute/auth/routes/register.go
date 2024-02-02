package routes

import (
	"context"
	"net/http"

	"practs/mainRoute/pb"

	"github.com/gin-gonic/gin"
)

type RegisterRequestBody struct {
	Logname  string `json:"logname"`
	Password string `json:"password"`
}

// @Summary      Registration
// @Tags         auth
// @Description  Registration new account
// @Produce      json
// @Param        main  body      pb.RegisterRequest  true  "Login and pass"
// @Success      201   {object}  pb.RegisterResponse
// @Failure      409   {object}  pb.RegisterResponse
// @Router       /auth/register [post]
func Register(ctx *gin.Context, c pb.AuthServiceClient) {
	body := RegisterRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Register(context.Background(), &pb.RegisterRequest{
		Logname:  body.Logname,
		Password: body.Password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(int(res.Status), &res)
}
