package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"reflect"
	"sync"
	"testing"
)

func TestNewTaskService(t *testing.T) {
	tests := []struct {
		name string
		want *TaskService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTaskService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTaskService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTaskService_createTask(t *testing.T) {
	type fields struct {
		tasks      map[int]Task
		nextTaskID int
		mu         sync.Mutex
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TaskService{
				tasks:      tt.fields.tasks,
				nextTaskID: tt.fields.nextTaskID,
				mu:         tt.fields.mu,
			}
			s.createTask(tt.args.w, tt.args.r)
		})
	}
}

func TestTaskService_deleteTask(t *testing.T) {
	type fields struct {
		tasks      map[int]Task
		nextTaskID int
		mu         sync.Mutex
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TaskService{
				tasks:      tt.fields.tasks,
				nextTaskID: tt.fields.nextTaskID,
				mu:         tt.fields.mu,
			}
			s.deleteTask(tt.args.w, tt.args.r)
		})
	}
}

func TestTaskService_getTask(t *testing.T) {
	type fields struct {
		tasks      map[int]Task
		nextTaskID int
		mu         sync.Mutex
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TaskService{
				tasks:      tt.fields.tasks,
				nextTaskID: tt.fields.nextTaskID,
				mu:         tt.fields.mu,
			}
			s.getTask(tt.args.w, tt.args.r)
		})
	}
}

func TestTaskService_listTasks(t *testing.T) {
	type fields struct {
		tasks      map[int]Task
		nextTaskID int
		mu         sync.Mutex
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TaskService{
				tasks:      tt.fields.tasks,
				nextTaskID: tt.fields.nextTaskID,
				mu:         tt.fields.mu,
			}
			s.listTasks(tt.args.w, tt.args.r)
		})
	}
}

func TestTaskService_updateTask(t *testing.T) {
	type fields struct {
		tasks      map[int]Task
		nextTaskID int
		mu         sync.Mutex
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TaskService{
				tasks:      tt.fields.tasks,
				nextTaskID: tt.fields.nextTaskID,
				mu:         tt.fields.mu,
			}
			s.updateTask(tt.args.w, tt.args.r)
		})
	}
}

func Test_respondWithError(t *testing.T) {
	type args struct {
		w       http.ResponseWriter
		code    int
		message string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			respondWithError(tt.args.w, tt.args.code, tt.args.message)
		})
	}
}

func Test_respondWithJSON(t *testing.T) {
	type args struct {
		w       http.ResponseWriter
		code    int
		payload interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			respondWithJSON(tt.args.w, tt.args.code, tt.args.payload)
		})
	}
}

func Test_setupRouter(t *testing.T) {
	type args struct {
		service *TaskService
	}
	tests := []struct {
		name string
		args args
		want *chi.Mux
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setupRouter(tt.args.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setupRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}
