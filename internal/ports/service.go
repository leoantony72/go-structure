package ports

import "test/internal/model"

type TestService interface {
	GetData(id string) *model.User
}
