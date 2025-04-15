package services

import (
	"github.com/prabhjotaulakh159/doc-save/repositories"
	"github.com/prabhjotaulakh159/doc-save/types"
	"strings"
)

type UserService interface {
	CreateNewUser(username string, password string) error
}

type CrudUserService struct {
	UserRepository repositories.UserRepository
	EncryptionService EncryptionService
}

func (c *CrudUserService) CreateNewUser(username string, password string) error {
	_username := strings.TrimSpace(username)
	_password := strings.TrimSpace(password)

	const minUsernameLen int = 4
	const minPasswordLen int = 8		
	
	if len(_username) < minUsernameLen || len(_password) < minPasswordLen {
		return &types.ValidationError{Message: "username and password must be atleast 4 characters long and 8 characters long respectively"}	
	}
	
	if _username == _password {
		return &types.ValidationError{Message: "username and password cannot be equal"}
	}
	
	usernameTaken, err := c.UserRepository.CheckIfUserExists(_username)
	if err != nil {
		return &types.ServerError{
			Message: "error in checking if username is already taken",
			InternalError: err,		
		}	
	} 
	
	if usernameTaken {
		return &types.ValidationError{
			Message: "username is already taken",		
		}	
	}
	
	encryptedPassword, err := c.EncryptionService.EncryptPassword(_password)
	if err != nil {
		return &types.ServerError{
			Message: "error securing user credentials",
			InternalError: err,	
		}
	}
	
	if err := c.UserRepository.CreateNewUser(_username, encryptedPassword); err != nil{
		return &types.ServerError{
			Message: "error creating a new user",
			InternalError: err,	
		}
	} 
	
	return nil	
}