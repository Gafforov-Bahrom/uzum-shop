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
		return
	}

	c.JSON(200, dtoToProduct(product))
}

func (r *Router) GetProducts(c *gin.Context) {
	temp := &dto.ListProductIn{
		Limit:  100,
		Offset: 0,
		Query:  c.Query("query"),
	}

	products, err := r.productService.List(c, temp)
	if errors.Is(err, sql.ErrNoRows) {
		c.AbortWithStatus(404)
	}
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	res := GetProductsResponse{
		Count:    products.Count,
		Products: listToProduct(products.Items),
	}
	c.JSON(200, res)
}

func (r *Router) AddProduct(c *gin.Context) {
	var request AddProductRequest

	token := c.GetHeader("Authorization")
	userId, err := r.getUserId(c, token)
	if err != nil {
		c.AbortWithError(401, err)
		return
	}

	err = c.ShouldBindJSON(&request)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	id, err := r.basketService.AddProduct(c, &dto.Basket{
		ProductId: dto.TypeID(request.ProductId),
		UserId:    dto.TypeID(userId),
		Count:     request.Count,
	})
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.JSON(201, gin.H{
		"id": id,
	})
}

func (r *Router) UpdateProduct(c *gin.Context) {
	var request UpdateProductRequest

	token := c.GetHeader("Authorization")
	userId, err := r.getUserId(c, token)
	if err != nil {
		c.AbortWithError(401, err)
		return
	}
	err = c.ShouldBindJSON(&request)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	product, err := r.basketService.UpdateBasket(c, &dto.Basket{
		UserId:    dto.TypeID(userId),
		ProductId: dto.TypeID(request.ProductId),
		Count:     request.Count,
	})
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, gin.H{
		"count": product.Count,
	})
}

func (r *Router) DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	token := c.GetHeader("Authorization")
	userId, err := r.getUserId(c, token)
	if err != nil {
		c.AbortWithError(401, err)
		return
	}

	err = r.basketService.DeleteProduct(c, &dto.Basket{
		UserId:    dto.TypeID(userId),
		ProductId: dto.TypeID(id),
	})
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.JSON(204, gin.H{})

}

func (r *Router) GetBasket(c *gin.Context) {
	token := c.GetHeader("Authorization")
	userId, err := r.getUserId(c, token)
	if err != nil {
		c.AbortWithError(401, err)
		return
	}

	basket, err := r.basketService.ListBaskets(c, dto.TypeID(userId))
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, gin.H{
		"basket": basket,
	})
}

func (r *Router) CreateOrder(c *gin.Context) {
	token := c.GetHeader("Authorization")
	userId, err := r.getUserId(c, token)
	if err != nil {
		c.AbortWithError(401, err)
		return
	}

	order, err := r.orderService.CreateOrder(c, dto.TypeID(userId))
	if err != nil {
		c.AbortWithStatusJSON(404, gin.H{
			"message": "not enough products",
		})
		return
	}
	c.JSON(200, gin.H{
		"order": order,
	})
}

func (r *Router) CancelOrder(c *gin.Context) {
	token := c.GetHeader("Authorization")
	userId, err := r.getUserId(c, token)
	_ = userId
	if err != nil {
		c.AbortWithError(401, err)
	}

	orderId, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	err = r.orderService.DeleteOrder(c, dto.DeleteOrderRequest{
		Id:     dto.TypeID(orderId),
		UserId: dto.TypeID(userId),
	})
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(204, gin.H{})
}
