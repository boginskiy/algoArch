package handlers

import (
	"aggregatorProject/internal/converter"
	"aggregatorProject/internal/service"
)

type UserHandler struct {
	Service   service.UserService
	Converter converter.UserConverter
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// func(u *UserHandler)

// type Handler interface {
//     ServeHTTP(ResponseWriter, *Request)
// }

// TODO...
// Реализация Hendler далее ...
