package shop

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
	"github.com/gin-gonic/gin"
)

func (r *Router) GetProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	product, err := r.productService.GetProduct(c, dto.TypeID(id))
	if errors.Is(err, sql.ErrNoRows) {
		c.AbortWithStatus(404)
		return
	}
	if err != nil {
		c.AbortWithError(500, err)
	}

	c.JSON(200, dtoToProduct(product))
}

func (r *Router) AddProduct(c *gin.Context) {
	var request AddProductRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	id, err := r.basketService.AddProduct(c, &dto.Basket{
		ProductId: dto.TypeID(request.ProductId),
		UserId:    dto.TypeID(request.UserId),
		Count:     request.Count,
	})
	if err != nil {
		c.AbortWithError(500, err)
	}

	c.JSON(201, gin.H{
		"id": id,
	})
}
