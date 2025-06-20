package handler

import (
	"log"
	"net/http"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (uh *UserHandler) Profile(w http.ResponseWriter, r *http.Request) {
	c := CurrentUser(r)
	if c == nil {
		http.Error(w, "unauthenticated", http.StatusUnauthorized)
		return
	}
	_, err := w.Write([]byte("Hello, " + c.Issuer + "!"))
	if err != nil {
		log.Println("error writing response:", err)
	}
}
