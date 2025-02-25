package userrepo

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

type MockUserRepository struct {
	mock.Mock
}

func (repo *MockUserRepository) FindUserById(ctx context.Context, id string) (*User, error) {
	args := repo.Called(ctx, id)
	if args.Get(0) != nil {
		return args.Get(0).(*User), args.Error(1)
	}

	return nil, args.Error(1)
}

func TestUserService_getUser_MockTestify(t *testing.T) {
	// проинициализируем
	mockUserRepository := new(MockUserRepository)
	mockUserRepository.On("FindUserById", mock.Anything, "1").Return(&User{Id: "1", Name: "Denis", Email: "denis@example.com"}, nil)

	service := UserService{repo: mockUserRepository}
	user, err := service.GetUser(context.Background(), "1")
	require.NoError(t, err)
	require.Equal(t, "1", user.Id)
}
