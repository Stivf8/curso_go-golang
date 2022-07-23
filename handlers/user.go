package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/segmentio/ksuid"
	"stiven.com/go/rest-ws/models"
	"stiven.com/go/rest-ws/repository"
	"stiven.com/go/rest-ws/server"
)

type SignUpResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignUpHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request = SignUpRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		//asigna un numero random al id
		id, err := ksuid.NewRandom()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var user = models.User{
			Email:    request.Email,
			Password: request.Password,
			Id:       id.String(),
		}
		err = repository.InsertUser(r.Context(), &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//devolvemos en tipo json
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(SignUpResponse{
			ID:    user.Id,
			Email: user.Email,
		})

	}
}
