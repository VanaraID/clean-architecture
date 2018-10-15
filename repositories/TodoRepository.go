package repositories

import (
	"github.com/VanaraID/clean-architecture/interfaces"
	"github.com/VanaraID/clean-architecture/models"

	"fmt"
)

type TodoRepository struct {
	interfaces.IDBHandler
}

func (repository *TodoRepository) FetchTodo(id int) (models.Todo, error) {

	row, err := repository.Query(fmt.Sprintf("SELECT * FROM todos WHERE id = '%d'", id))
	if err != nil {
		return models.Todo{}, err
	}

	var todo models.Todo

	row.Next()
	row.Scan(&todo.Id, &todo.Todo, &todo.Done)

	return todo, nil
}

func (repository *TodoRepository) FetchAllTodo() ([]models.Todo, error) {

	row, err := repository.Query(fmt.Sprintf("SELECT * FROM todos"))
	if err != nil {
		return []models.Todo{}, err
	}

	var (
		todo  models.Todo
		todos []models.Todo
	)

	for row.Next() {
		row.Scan(&todo.Id, &todo.Todo, &todo.Done)

		todos = append(todos, todo)
	}

	return todos, nil
}
