package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/VanaraID/clean-architecture/models"
	"github.com/VanaraID/clean-architecture/payloads"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"

	"github.com/VanaraID/clean-architecture/interfaces/mocks"
)

func TestGetTodo(t *testing.T) {
	// create an instance of our test object
	todoService := new(mocks.ITodoService)

	todo := models.Todo{Id: 1, Todo: "Code", Done: 0}

	todoService.On("GetTodoByID", "1").Return(todo, nil)

	todoController := TodoController{todoService}

	// call the code we are testing
	req := httptest.NewRequest("GET", "http://localhost:6969/todo/1", nil)
	w := httptest.NewRecorder()

	r := chi.NewRouter()
	r.HandleFunc("/todo/{id}", todoController.GetTodo)

	r.ServeHTTP(w, req)

	expectedResult := payloads.SuccessGetTodo{Code: http.StatusOK, Response: todo}

	actualResult := payloads.SuccessGetTodo{}

	json.NewDecoder(w.Body).Decode(&actualResult)
	// assert that the expectations were met
	assert.Equal(t, expectedResult, actualResult)
}

func TestGetAllTodo(t *testing.T) {
	// create an instance of our test object
	todoService := new(mocks.ITodoService)

	var todos = []models.Todo{}
	todos = append(todos, models.Todo{Id: 1, Todo: "Eat", Done: 1})
	todos = append(todos, models.Todo{Id: 2, Todo: "Code", Done: 0})
	todos = append(todos, models.Todo{Id: 3, Todo: "Sleep", Done: 0})

	todoService.On("GetTodos").Return(todos, nil)

	todoController := TodoController{todoService}

	// call the code we are testing
	req := httptest.NewRequest("GET", "http://localhost:6969/todos", nil)
	w := httptest.NewRecorder()

	r := chi.NewRouter()
	r.HandleFunc("/todos", todoController.GetAllTodo)

	r.ServeHTTP(w, req)

	expectedResult := payloads.SuccessGetAllTodo{Code: http.StatusOK, Response: todos}

	actualResult := payloads.SuccessGetAllTodo{}

	json.NewDecoder(w.Body).Decode(&actualResult)
	// assert that the expectations were met
	assert.Equal(t, expectedResult, actualResult)
}
