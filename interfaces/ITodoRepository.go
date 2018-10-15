package interfaces

import (
	"github.com/VanaraID/clean-architecture/models"
)

type ITodoRepository interface {
	FetchTodo(id int) (models.Todo, error)
	FetchAllTodo() ([]models.Todo, error)
}
