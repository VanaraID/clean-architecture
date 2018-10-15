package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/VanaraID/clean-architecture/interfaces"
	"github.com/VanaraID/clean-architecture/payloads"
	"github.com/go-chi/chi"
)

type TodoController struct {
	interfaces.ITodoService
}

func (controller *TodoController) GetTodo(res http.ResponseWriter, req *http.Request) {

	id := chi.URLParam(req, "id")

	todo, err := controller.GetTodoByID(id)
	if err != nil {
		//Handle error
	}

	json.NewEncoder(res).Encode(payloads.Success{Code: http.StatusOK, Response: todo})
}

func (controller *TodoController) GetAllTodo(res http.ResponseWriter, req *http.Request) {

	todos, err := controller.GetTodos()
	if err != nil {
		//Handle error
	}

	json.NewEncoder(res).Encode(payloads.Success{Code: http.StatusOK, Response: todos})
}
