package basket

type Service struct {
	repo BasketRepository
}

func NewService(repo BasketRepository) *Service {
	return &Service{
		repo: repo,
	}
}
