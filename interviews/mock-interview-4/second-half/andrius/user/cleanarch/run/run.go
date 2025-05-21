package run

import (
	"cleanarch/internal/controller"
	"cleanarch/internal/repository"
	"cleanarch/internal/usecase"
	"log"
	"net/http"
)

func Run() {
	// dependency injection
	repo := repository.NewEmpRepo()
	uc := usecase.NewEmpUseCase(repo)
	router := controller.NewEmpRouter(uc)

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
