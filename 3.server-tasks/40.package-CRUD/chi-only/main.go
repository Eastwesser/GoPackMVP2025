package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/go-chi/chi/v5"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type TaskService struct {
	tasks      map[int]Task
	nextTaskID int
	mu         sync.Mutex
}

func NewTaskService() *TaskService {
	return &TaskService{
		tasks:      make(map[int]Task),
		nextTaskID: 1,
	}
}

func main() {
	taskService := NewTaskService()
	r := setupRouter(taskService)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func setupRouter(service *TaskService) *chi.Mux {
	r := chi.NewRouter()

	r.Route("/tasks", func(r chi.Router) {
		r.Get("/", service.listTasks)         // GET /tasks - список всех задач
		r.Post("/", service.createTask)       // POST /tasks - создать новую задачу
		r.Get("/{id}", service.getTask)       // GET /tasks/{id} - получить задачу по ID
		r.Put("/{id}", service.updateTask)    // PUT /tasks/{id} - обновить задачу
		r.Delete("/{id}", service.deleteTask) // DELETE /tasks/{id} - удалить задачу
	})

	return r
}

func (s *TaskService) listTasks(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	defer s.mu.Unlock()

	taskList := make([]Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		taskList = append(taskList, task)
	}

	respondWithJSON(w, http.StatusOK, taskList)
}

func (s *TaskService) createTask(w http.ResponseWriter, r *http.Request) {
	var newTask Task
	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	newTask.ID = s.nextTaskID
	s.tasks[newTask.ID] = newTask
	s.nextTaskID++

	respondWithJSON(w, http.StatusCreated, newTask)
}

func (s *TaskService) getTask(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid task-for-understanding ID")
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	task, exists := s.tasks[id]
	if !exists {
		respondWithError(w, http.StatusNotFound, "Task not found")
		return
	}

	respondWithJSON(w, http.StatusOK, task)
}

func (s *TaskService) updateTask(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid task-for-understanding ID")
		return
	}

	var updatedTask Task
	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.tasks[id]; !exists {
		respondWithError(w, http.StatusNotFound, "Task not found")
		return
	}

	updatedTask.ID = id
	s.tasks[id] = updatedTask

	respondWithJSON(w, http.StatusOK, updatedTask)
}

func (s *TaskService) deleteTask(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid task-for-understanding ID")
		return
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.tasks[id]; !exists {
		respondWithError(w, http.StatusNotFound, "Task not found")
		return
	}

	delete(s.tasks, id)
	w.WriteHeader(http.StatusNoContent)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error marshalling JSON"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
