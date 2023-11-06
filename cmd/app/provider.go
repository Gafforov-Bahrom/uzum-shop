package main

import (
	"github.com/Gafforov-Bahrom/uzum_shop/internal/config"
	"github.com/jmoiron/sqlx"
)

type serviceProvider struct {
	cfg *config.Config
	db  *sqlx.DB
}

func NewServiceProvider(cfg *config.Config) *serviceProvider {
	return &serviceProvider{cfg: cfg}
}
