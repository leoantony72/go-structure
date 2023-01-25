package services

import "test/internal/model"

func (t *TestService) GetData(id string) *model.User {
	user := t.repo.GetData(id)
	return user
}
