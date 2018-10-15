package main

import (
	"database/sql"
	"sync"

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

	sqlConn, _ := sql.Open("mysql", "root:root@/vanara")
	MySQLHandler := &infrastructures.MySQLHandler{}
	MySQLHandler.Conn = sqlConn

	todoRepository := &repositories.TodoRepository{MySQLHandler}
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
