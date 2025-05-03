package main

import (
	"andrius-cleanarch-map/internal/infrastructure"
	"andrius-cleanarch-map/internal/interfaces"
	"andrius-cleanarch-map/internal/usecase"
	"fmt"
)

func main() {
	repo := infrastructure.NewInMemoryUserRepo()
	interactor := usecase.NewUserInteractor(repo)
	handler := interfaces.NewHandler(interactor)

	handler.CreateUser("1", "Alice")
	handler.CreateUser("2", "Bob")

	fmt.Println(handler.GetAllUsers())
	fmt.Println(handler.GetUserByID("1"))
}
