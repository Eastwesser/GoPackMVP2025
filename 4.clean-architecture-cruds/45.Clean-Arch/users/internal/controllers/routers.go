package controllers

import (
	"cleanuser/internal/usecase"
	"github.com/go-chi/chi/v5"
)

func SetupUserRouter(uc *usecase.UserUseCase) *chi.Mux {
	// Настройка роутера
	r := chi.NewRouter()
	handler := NewUserHandler(uc)
	// CRUDs
	r.Route("/user", func(r chi.Router) {
		// POST /users - создать нового пользователя
		r.Post("/", handler.CreateUser)
		//r.Method("POST", "/", http.HandlerFunc(handler.Create))

		// GET /users/{id} - получить пользователя по ID
		r.Get("/{id}", handler.ReadUser)
		//r.Method("GET", "/{id}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//	// Извлекаем ID из контекста
		//	id := chi.URLParam(r, "id")
		//	handler.Get(w, r.WithContext(context.WithValue(r.Context(), "id", id)))
		//}))

		// PUT /users/{id} - обновить пользователя
		//r.Put("/{id}", handler.Update)
		//r.Method("PUT", "/{id}", http.HandlerFunc(handler.Update))

		// DELETE /users/{id} - удалить пользователя
		//r.Delete("/{id}", handler.Delete)
		//r.Method("GET", "/{id}", http.HandlerFunc(handler.Get))

		//// GET /users - список всех пользователей
		//r.Get("/", handler.GetAll)
		////r.Method("GET", "/", http.HandlerFunc(handler.GetAll))
	})
	return r
}
