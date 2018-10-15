package interfaces

import "github.com/VanaraID/clean-architecture/models"

type ITodoService interface {
	GetTodoByID(id string) (models.Todo, error)
	GetTodos() ([]models.Todo, error)
}
