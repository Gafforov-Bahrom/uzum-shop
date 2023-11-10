package user

import (
	"context"

	"github.com/Gafforov-Bahrom/uzum_shop/internal/dto"
	desc "github.com/Gafforov-Bahrom/uzum_shop/pkg/grpc/shop"
)

func (s *grpcService) ListUsers(ctx context.Context, in *desc.ListUsersRequest) (*desc.ListUsersResponse, error) {
	dtoList := make([]dto.TypeID, len(in.UserIds))
	for i, v := range in.UserIds {
		dtoList[i] = dto.TypeID(v)
	}

	out, err := s.userService.ListUsers(ctx, dtoList)
	if err != nil {
		return nil, err
	}
	return &desc.ListUsersResponse{Items: listToDesc(out)}, nil
}
