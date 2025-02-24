package userrepo

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestUserRepositoryImpl_GetUser_Mock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создаем мок репозитория
	mockRepo := NewMockUserRepository(ctrl) // Используем напрямую из `userrepo`
	mockRepo.EXPECT().
		FindById(gomock.Any(), "1").
		Return(&User{Id: "1", Name: "Denis", Email: "denis@gmail.com"}, nil)

	// Создаем сервис с мок-репозиторием
	service := &UserRepositoryImpl{repo: mockRepo}

	// Вызываем метод и проверяем результат
	user, err := service.GetUser(context.Background(), "1")
	require.NoError(t, err)
	require.Equal(t, "1", user.Id)
	require.Equal(t, "Denis", user.Name)
	require.Equal(t, "denis@gmail.com", user.Email)
}
