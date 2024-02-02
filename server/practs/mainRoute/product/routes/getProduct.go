package routes

import (
	"context"
	"net/http"

	"practs/mainRoute/pb"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Summary      Get Product
// @Tags         product
// @Description  Отправка каталога магазина
// @Produce      json
// @Success      200   {object}  pb.GetProductResponse
// @Failure      404   {object}  pb.GetProductResponse
// @Router       /product [get]
func GetProduct(ctx *gin.Context, c pb.ProductServiceClient) {

	res, err := c.GetProduct(context.Background(), &pb.Empty{})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
