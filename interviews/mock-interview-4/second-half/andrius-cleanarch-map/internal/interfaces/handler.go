package interfaces

import (
	"andrius-cleanarch-map/internal/usecase"
	"fmt"
)

// controllers/routers

type Handler struct {
	interactor *usecase.UserInteractor
}

func NewHandler(i *usecase.UserInteractor) *Handler {
	return &Handler{interactor: i}
}

func (h *Handler) CreateUser(id, name string) {
	h.interactor.CreateUser(id, name)
}

func (h *Handler) GetAllUsers() string {
	users := h.interactor.GetAll()
	return fmt.Sprintf("Users: %+v", users)
}

func (h *Handler) GetUserByID(id string) string {
	user, found := h.interactor.GetByID(id)
	if !found {
		return "User not found"
	}
	return fmt.Sprintf("User: %+v", user)
}
