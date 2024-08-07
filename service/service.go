package service

import (
	"github.com/itzaddddd/ticket-reserve/repository"
)

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}
