package controllers

import (
	"net/http"
)

type UserController interface {
	CreateNewUser(writer http.ResponseWriter, request *http.Request)
}

type CrudUserController struct{}

func (c *CrudUserController) CreateNewUser(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
}