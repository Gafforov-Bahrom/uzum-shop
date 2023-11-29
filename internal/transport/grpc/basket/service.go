package basket

import (
	"context"

	"github.com/Gafforov-Bahrom/uzum_shop/pkg/grpc/login_v1"
	desc "github.com/Gafforov-Bahrom/uzum_shop/pkg/grpc/shop"
)

type grpcService struct {
	desc.UnimplementedBasketServiceServer
	basketService BasketService
	loginClient   login_v1.LoginV1Client
}

func NewService(basketService BasketService, loginClient login_v1.LoginV1Client) desc.BasketServiceServer {
	return &grpcService{
		basketService: basketService,
		loginClient:   loginClient,
	}
}

func (s *grpcService) getUserId(ctx context.Context, token string) (uint64, error) {
	out, err := s.loginClient.GetUserId(ctx, &login_v1.GetUserIdRequest{
		AccessToken: token,
	})
	if err != nil {
		return 0, err
	}
	return out.UserId, nil
}
