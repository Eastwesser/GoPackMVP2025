package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestUserRepositoryImpl_GetUser_MockTestify(t *testing.T) {
	mockRepo := new(MockUserRepository)

	// Ожидаем вызов FindById с любым контекстом и ID "1", возвращаем тестового пользователя
	mockRepo.On("FindById", mock.Anything, "1").Return(&User{Id: "1", Name: "Denis", Email: "denis@example.com"}, nil)

	// Создаём сервис с мок-репозиторием
	service := UserRepositoryImpl{repo: mockRepo}

	// Вызываем метод
	user, err := service.GetUser(context.Background(), "1")

	// Проверяем, что всё сработало
	require.NoError(t, err)
	require.Equal(t, "1", user.Id)
	require.Equal(t, "Denis", user.Name)
	require.Equal(t, "denis@example.com", user.Email)

	// Проверяем, что мок реально был вызван
	mockRepo.AssertExpectations(t)
}
