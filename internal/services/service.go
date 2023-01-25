package services

import "test/internal/ports"

type TestService struct {
	repo ports.TestRepo
}

func NewTestService(repo ports.TestRepo) ports.TestService {
	return &TestService{
		repo: repo,
	}
}
