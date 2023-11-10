package user

import desc "github.com/Gafforov-Bahrom/uzum_shop/pkg/grpc/shop"

type grpcService struct {
	desc.UnimplementedUserServiceServer
	userService UserService
}

func NewService(userService UserService) desc.UserServiceServer {
	return &grpcService{
		userService: userService,
	}
}
