package order

type Service struct {
	repo       OrderRepository
	userRepo   UserRepository
	basketRepo BasketRepository
}

func NewService(repo OrderRepository, userRepo UserRepository, basketRepo BasketRepository) *Service {
	return &Service{
		repo:       repo,
		userRepo:   userRepo,
		basketRepo: basketRepo,
	}
}

//update basket set order id
