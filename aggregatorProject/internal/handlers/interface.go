package handlers

import "net/http"

type UserHandler interface {
	Read(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
}
