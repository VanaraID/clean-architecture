package services

import (
	"strconv"

	"github.com/VanaraID/clean-architecture/interfaces"
	"github.com/VanaraID/clean-architecture/models"
)

type TodoService struct {
	interfaces.ITodoRepository
}

func (service *TodoService) GetTodoByID(id string) (models.Todo, error) {

	idInt, errAtoi := strconv.Atoi(id)
	if errAtoi != nil {
		// handle error
	}

	todo, err := service.FetchTodo(idInt)
	if err != nil {
		//Handle error
	}

	return todo, nil
}

func (service *TodoService) GetTodos() ([]models.Todo, error) {

	todos, err := service.FetchAllTodo()
	if err != nil {
		//Handle error
	}

	return todos, nil
}
