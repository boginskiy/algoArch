package converter

import (
	"aggregatorProject/internal/model"
	"net/http"
)

type UserConverter interface {
	ConvertBytesToUser(*http.Request) (*model.User, error)
	ConvertUserToBytes(*model.User) ([]byte, error)
}
