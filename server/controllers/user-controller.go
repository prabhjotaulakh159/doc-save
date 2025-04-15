package controllers

import (
	"net/http"
	"github.com/prabhjotaulakh159/doc-save/services"
	"github.com/prabhjotaulakh159/doc-save/types"
	"encoding/json"
	"errors"
	"log"
)

type usernamePasswordRequest struct {
	Username string `json:"username", required: "true"`
	Password string `json:"password", required: "true"`
}

type UserController interface {
	CreateNewUser(w http.ResponseWriter, r *http.Request)
}

type CrudUserController struct{
	UserService services.UserService
}

func (c *CrudUserController) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	var req usernamePasswordRequest
	dec := json.NewDecoder(r.Body)
   dec.DisallowUnknownFields()
   err := dec.Decode(&req)
   if err != nil {
		http.Error(w, "only username and password are accepted fields", http.StatusBadRequest)   
		return
   }
	if err := c.UserService.CreateNewUser(req.Username, req.Password); err != nil {
		var vError *types.ValidationError
		var dError *types.DbError
		if (errors.As(err, &vError)) {
			http.Error(w, vError.Message, http.StatusBadRequest)	
			return	
		} 	
		if (errors.As(err, &dError)) {
			log.Println(err)
			http.Error(w, dError.Message, http.StatusInternalServerError)
			return		
		}
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}