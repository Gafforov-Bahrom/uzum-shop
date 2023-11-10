package order

type Service struct {
	repo     OrderRepository
	userRepo UserRepository
}

func NewService(repo OrderRepository, userRepo UserRepository) *Service {
	return &Service{
		repo:     repo,
		userRepo: userRepo,
	}
}

//update basket set order id
