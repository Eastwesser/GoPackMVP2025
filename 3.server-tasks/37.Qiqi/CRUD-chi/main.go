package main

import (
	"encoding/json" // Для работы с JSON (кодирование/декодирование)
	"net/http"      // HTTP сервер и клиент
	"strconv"       // Преобразование строк в числа и обратно
	"sync"          // Для синхронизации (мьютексы)

	"github.com/go-chi/chi/v5" // Роутер для HTTP запросов
)

// User определяет структуру данных пользователя.
// Теги `json:` используются для указания имен полей при сериализации в JSON.
type User struct {
	ID    int    `json:"id"`    // Уникальный идентификатор
	Name  string `json:"name"`  // Имя пользователя
	Email string `json:"email"` // Электронная почта
}

// Глобальные переменные для хранения состояния приложения
var (
	// users - in-memory хранилище пользователей в виде map (ключ - ID пользователя)
	users = make(map[int]User)

	// usersMutex - RWMutex для безопасного доступа к users из разных горутин
	// RWMutex позволяет нескольким читателям или одному писателю
	usersMutex sync.RWMutex

	// idCounter - счетчик для генерации новых ID пользователей
	idCounter = 1
)

func main() {
	// Создаем новый роутер с помощью chi
	r := chi.NewRouter()

	// Группируем маршруты для пользователей под префиксом /users
	// Это позволяет создавать вложенные маршруты
	r.Route("/users", func(r chi.Router) {
		// GET /users - получение списка всех пользователей
		r.Get("/", listUsers)

		// POST /users - создание нового пользователя
		r.Post("/", createUser)

		// Вложенная группа маршрутов для операций с конкретным пользователем
		// {id} - параметр маршрута, который будет доступен через chi.URLParam
		r.Route("/{id}", func(r chi.Router) {
			// GET /users/{id} - получение пользователя по ID
			r.Get("/", getUser)

			// PUT /users/{id} - обновление пользователя
			r.Put("/", updateUser)

			// DELETE /users/{id} - удаление пользователя
			r.Delete("/", deleteUser)
		})
	})

	// Запускаем HTTP сервер на порту 8080
	// Вторым аргументом передаем наш роутер, который будет обрабатывать все запросы
	http.ListenAndServe(":8080", r)
}

// listUsers обрабатывает GET запрос для получения списка всех пользователей
func listUsers(w http.ResponseWriter, r *http.Request) {
	// Блокируем для чтения (могут быть другие читатели одновременно)
	usersMutex.RLock()
	// Гарантируем, что мьютекс будет разблокирован при выходе из функции
	defer usersMutex.RUnlock()

	// Создаем слайс для хранения пользователей с начальной емкостью = размеру map
	userList := make([]User, 0, len(users))

	// Итерируемся по map и добавляем пользователей в слайс
	for _, user := range users {
		userList = append(userList, user)
	}

	// Отправляем JSON ответ со списком пользователей
	respondWithJSON(w, http.StatusOK, userList)
}

// createUser обрабатывает POST запрос для создания нового пользователя
func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser User

	// Декодируем тело запроса (JSON) в структуру User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		// Если ошибка декодирования - возвращаем HTTP 400
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Блокируем для записи (никакие другие операции не могут выполняться)
	usersMutex.Lock()
	// Гарантируем разблокировку
	defer usersMutex.Unlock()

	// Устанавливаем ID пользователя и увеличиваем счетчик
	newUser.ID = idCounter
	users[newUser.ID] = newUser
	idCounter++

	// Возвращаем созданного пользователя с HTTP статусом 201 (Created)
	respondWithJSON(w, http.StatusCreated, newUser)
}

// getUser обрабатывает GET запрос для получения пользователя по ID
func getUser(w http.ResponseWriter, r *http.Request) {
	// Получаем параметр id из URL (например, /users/123)
	idStr := chi.URLParam(r, "id")

	// Преобразуем строку в число
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// Если ID не число - возвращаем ошибку
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	// Блокируем для чтения
	usersMutex.RLock()
	defer usersMutex.RUnlock()

	// Пытаемся найти пользователя в map
	user, exists := users[id]
	if !exists {
		// Если не найден - возвращаем HTTP 404
		respondWithError(w, http.StatusNotFound, "User not found")
		return
	}

	// Возвращаем найденного пользователя
	respondWithJSON(w, http.StatusOK, user)
}

// updateUser обрабатывает PUT запрос для обновления пользователя
func updateUser(w http.ResponseWriter, r *http.Request) {
	// Получаем ID из URL
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var updatedUser User
	// Декодируем тело запроса
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Блокируем для записи
	usersMutex.Lock()
	defer usersMutex.Unlock()

	// Проверяем существование пользователя
	if _, exists := users[id]; !exists {
		respondWithError(w, http.StatusNotFound, "User not found")
		return
	}

	// Обновляем данные пользователя (сохраняем ID из URL)
	updatedUser.ID = id
	users[id] = updatedUser

	// Возвращаем обновленного пользователя
	respondWithJSON(w, http.StatusOK, updatedUser)
}

// deleteUser обрабатывает DELETE запрос для удаления пользователя
func deleteUser(w http.ResponseWriter, r *http.Request) {
	// Получаем ID из URL
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	// Блокируем для записи
	usersMutex.Lock()
	defer usersMutex.Unlock()

	// Проверяем существование пользователя
	if _, exists := users[id]; !exists {
		respondWithError(w, http.StatusNotFound, "User not found")
		return
	}

	// Удаляем пользователя из map
	delete(users, id)

	// Возвращаем HTTP 204 (No Content) - успешное удаление без тела ответа
	w.WriteHeader(http.StatusNoContent)
}

// respondWithError создает JSON ответ с ошибкой
func respondWithError(w http.ResponseWriter, code int, message string) {
	// Используем respondWithJSON, передавая map с ключом "error"
	respondWithJSON(w, code, map[string]string{"error": message})
}

// respondWithJSON создает JSON ответ с произвольными данными
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	// Кодируем данные в JSON (игнорируем ошибку для простоты)
	response, _ := json.Marshal(payload)

	// Устанавливаем заголовки
	w.Header().Set("Content-Type", "application/json")
	// Устанавливаем HTTP статус код
	w.WriteHeader(code)
	// Пишем тело ответа
	w.Write(response)
}
