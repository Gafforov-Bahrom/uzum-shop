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
	}

	c.JSON(201, gin.H{
		"id": id,
	})
}

func (r *Router) UpdateProduct(c *gin.Context) {
	var request UpdateProductRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	product, err := r.basketService.UpdateBasket(c, &dto.Basket{
		UserId:    dto.TypeID(request.UserId),
		ProductId: dto.TypeID(request.ProductId),
		Count:     request.Count,
	})
	if err != nil {
		c.AbortWithError(500, err)
	}
	c.JSON(200, gin.H{
		"count": product.Count,
	})
}

func (r *Router) DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		c.AbortWithError(400, err)
	}

	var request DeleteProductRequest
	err = c.ShouldBindJSON(&request)
	if err != nil {
		c.AbortWithError(400, err)
	}

	err = r.basketService.DeleteProduct(c, &dto.Basket{
		UserId:    dto.TypeID(request.UserId),
		ProductId: dto.TypeID(id),
	})
	if err != nil {
		c.AbortWithError(500, err)
	}

	c.JSON(200, "product deleted")

}

func (r *Router) GetBasket(c *gin.Context) {
	var request GetBasketRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.AbortWithError(400, err)
	}

	basket, err := r.basketService.ListBaskets(c, dto.TypeID(request.UserId))
	if err != nil {
		c.AbortWithError(500, err)
	}
	c.JSON(200, gin.H{
		"basket": basket,
	})
}

func (r *Router) CreateOrder(c *gin.Context) {
	var request CreateOrderRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.AbortWithError(400, err)
	}

	order, err := r.orderService.CreateOrder(c, dto.TypeID(request.UserId))
	if err != nil {
		c.AbortWithError(500, err)
	}
	c.JSON(200, gin.H{
		"order": order,
	})
}

func (r *Router) CancelOrder(c *gin.Context) {
	orderId, err := strconv.Atoi(c.Param("order_id"))
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	err = r.orderService.DeleteOrder(c, dto.TypeID(orderId))
	if err != nil {
		c.AbortWithError(500, err)
	}
	c.JSON(200, "order is deleted")
}
