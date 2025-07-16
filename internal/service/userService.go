package service

import (
	"go-ecommerce-app/internal/domain"
	"go-ecommerce-app/internal/dto"
	"log"
)
type UserService struct {

}

func (s UserService) FindUserByEmail(email string)(*domain.User, error){
	// perform some db operation and business logic
	return nil, nil
}


func (s UserService) Signup(input dto.UserSignup) (string, error) {

	log.Println(input)

	return "this is my token as of now", nil
}

func (s UserService) findUserByEmail(email string) (*domain.User, error) {
	//perform some db operation
	//business logic

	return nil, nil
}

func (s UserService) Login(email string, password string) (string, error) {

	// generate token
	return "", nil
}
