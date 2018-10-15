package payloads

import "github.com/VanaraID/clean-architecture/models"

type SuccessGetTodo struct {
	Code     int
	Response models.Todo
}

type SuccessGetAllTodo struct {
	Code     int
	Response []models.Todo
}
