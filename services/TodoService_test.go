package services

import (
	"testing"

	"github.com/VanaraID/clean-architecture/interfaces/mocks"
	"github.com/VanaraID/clean-architecture/models"
	"github.com/stretchr/testify/assert"
)

func TestGetTodoByID(t *testing.T) {
	todoRepository := new(mocks.ITodoRepository)

	todo := models.Todo{}
	todo.Id = 1
	todo.Todo = "Code"
	todo.Done = 1

	todoRepository.On("FetchTodo", 1).Return(todo, nil)

	todoService := TodoService{todoRepository}

	expectedResult := todo

	actualResult, _ := todoService.FetchTodo(1)

	assert.Equal(t, expectedResult, actualResult)

}

func TestGetTodos(t *testing.T) {
	todoRepository := new(mocks.ITodoRepository)

	var todos = []models.Todo{}
	todos = append(todos, models.Todo{Id: 1, Todo: "Eat", Done: 1})
	todos = append(todos, models.Todo{Id: 2, Todo: "Code", Done: 0})
	todos = append(todos, models.Todo{Id: 3, Todo: "Sleep", Done: 0})

	todoRepository.On("FetchAllTodo").Return(todos, nil)

	todoService := TodoService{todoRepository}

	expectedResult := todos

	actualResult, _ := todoService.FetchAllTodo()

	assert.Equal(t, expectedResult, actualResult)

}
