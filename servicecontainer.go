package main

import (
	"sync"

	"database/sql"

	"github.com/VanaraID/clean-architecture/controllers"
	"github.com/VanaraID/clean-architecture/infrastructures"
	"github.com/VanaraID/clean-architecture/repositories"
	"github.com/VanaraID/clean-architecture/services"
)

type IServiceContainer interface {
	InjectTodoController() controllers.TodoController
}

type kernel struct{}

func (k *kernel) InjectTodoController() controllers.TodoController {

	sqlConn, _ := sql.Open("sqlite3", "/var/tmp/todos.db")
	sqliteHandler := &infrastructures.SQLiteHandler{}
	sqliteHandler.Conn = sqlConn

	todoRepository := &repositories.TodoRepository{sqliteHandler}
	todoService := &services.TodoService{todoRepository}
	todoController := controllers.TodoController{todoService}

	return todoController
}

var (
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}
