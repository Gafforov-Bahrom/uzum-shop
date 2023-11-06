package product

type Service struct {
	repo ProductRepository
}

func NewService(repo ProductRepository) *Service {
	return &Service{
		repo: repo,
	}
}
