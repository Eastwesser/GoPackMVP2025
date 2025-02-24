package main

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserService_GetUser(t *testing.T) {
	// бд у нас своя
	mockRepo := &MockUserRepository{
		mockData: map[string]*User{
			"1": {Id: "1", Name: "Denis", Email: "denis@example.com"},
		},
	}
	// тестируем GetUser
	service := &UserRepositoryImpl{repo: mockRepo} // вместо реального интерфейса мы подставляем мок, в котором лежит просто map (mockData)
	user, err := service.GetUser(context.Background(), "1")
	require.NoError(t, err)
	require.Equal(t, "1", user.Id)
}
