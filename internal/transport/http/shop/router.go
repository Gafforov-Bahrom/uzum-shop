package shop

import (
	"context"

	"github.com/Gafforov-Bahrom/uzum_shop/pkg/grpc/login_v1"
)

type Router struct {
	productService ProductService
	basketService  BasketService
	orderService   OrderService
	loginClient    login_v1.LoginV1Client
}

func NewRouter(productService ProductService, basketService BasketService, orderService OrderService, loginClient login_v1.LoginV1Client) *Router {
	return &Router{
		productService: productService,
		basketService:  basketService,
		orderService:   orderService,
		loginClient:    loginClient,
	}
}

func (s *Router) getUserId(ctx context.Context, token string) (uint64, error) {
	out, err := s.loginClient.GetUserId(ctx, &login_v1.GetUserIdRequest{
		AccessToken: token,
	})
	if err != nil {
		return 0, err
	}
	return out.UserId, nil
}
