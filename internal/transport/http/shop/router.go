package shop

type Router struct {
	productService ProductService
	basketService  BasketService
	orderService   OrderService
}

func NewRouter(productService ProductService, basketService BasketService, orderService OrderService) *Router {
	return &Router{
		productService: productService,
		basketService:  basketService,
		orderService:   orderService,
	}
}
