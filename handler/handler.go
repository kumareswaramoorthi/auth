package handler

import (
	"auth/database"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Handler struct {
	DB *sql.DB
}
type data map[string]interface{}

func JSONWriter(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/home" {
		fmt.Fprintf(w, "Welcome!")
		return
	}
}

func (h *Handler) FindUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		JSONWriter(w, data{
			"Error": "Method Not Allowed",
		}, http.StatusMethodNotAllowed)
		return
	}

	userName := r.Header.Get("username")
	if userName == "" {
		JSONWriter(w, data{
			"Error": "no username found in header",
		}, http.StatusInternalServerError)
		return
	}
	getUser, err := database.FindUserByUsername(userName, h.DB)
	if err != nil {
		JSONWriter(w, data{
			"Error": "internal server error",
		}, http.StatusNotFound)
		return
	}

	if getUser.Username == "" {
		JSONWriter(w, data{
			"Response": "User does not exist",
		}, http.StatusUnauthorized)
		return
	} else {
		JSONWriter(w, data{
			"Response": "User exists",
		}, http.StatusOK)
		return
	}

}
