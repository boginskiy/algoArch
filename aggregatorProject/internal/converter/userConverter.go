package converter

import (
	"aggregatorProject/internal/model"
	"net/http"
)

type UserConverter interface {
	ConvertDataToUser(http.Request) (*model.User, error)
}

type UserConvert struct {
}

func NewUserConvert() *UserConvert {
	return &UserConvert{}
}

func (c *UserConvert) ConvertDataToUser(req http.Request) (*model.User, error) {
	return &model.User{ID: 1, Name: "Dima"}, nil
}
