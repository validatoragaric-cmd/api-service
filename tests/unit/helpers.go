package main

import (
	"crypto/rand"
	"encoding/base32"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func generateRandomString() string {
	const (
		length  = 32
		charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	)

	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		log.Println(err)
	}
	for i, b := range b {
		b[i] = charset[b%byte(len(charset))]
	}
	return base32.StdEncoding.EncodeToString(b)
}

func generateToken(userId string) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"exp": time.Now().Add(time.Hour * 72).Unix(),
			"iat": time.Now().Unix(),
			"user_id": userId,
		},
	)
	return token.SignedString([]byte("secret_key"))
}

func generateResponse(data interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(data)
	}
}

func errorHandler(err error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func createRouter() *mux.Router {
	router := mux.NewRouter()
	return router
}

func configureRoutes(router *mux.Router) {
	router.HandleFunc("/users", get_users).Methods("GET")
	router.HandleFunc("/users/{id}", get_user).Methods("GET")
	router.HandleFunc("/users", create_user).Methods("POST")
	router.HandleFunc("/users/{id}", update_user).Methods("PUT")
	router.HandleFunc("/users/{id}", delete_user).Methods("DELETE")
}

func logRequest(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL.Path, r.RemoteAddr)
}