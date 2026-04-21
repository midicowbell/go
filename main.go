package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UserDTO struct {
	Username string `json:"username"`
	Age      int    `json:"age"`
	Is_admin bool   `json:"is_admin"`
}

var user UserDTO

func registerHandler(w http.ResponseWriter, r *http.Request) {

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Println("failed read json to struct")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	ans, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(ans); err != nil {
		fmt.Println("err: ", err)
		return
	}

}

func main() {
	http.HandleFunc("/register", registerHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("some error: ", err)
	}

}
