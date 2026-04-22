package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type DTO struct {
	Numbers []int `json:"numbers"`
}

type JsonResponse struct {
	Result  int       `json:"result"`
	TimeNow time.Time `json:"time"`
}

func calculateHande(w http.ResponseWriter, r *http.Request) {
	var data DTO
	actionA := r.URL.Query().Get("action")

	// Декодируем тело запроса
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest) // Лучше 400, если ошибка в данных пользователя
		w.Write([]byte(err.Error()))
		return
	}

	var resultAction int
	switch actionA {
	case "add":
		for _, val := range data.Numbers {
			resultAction += val
		}
	case "sub":
		// Если это вычитание, логично начинать с первого числа,
		// но здесь реализовано вычитание всех элементов из нуля.
		for _, val := range data.Numbers {
			resultAction -= val
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid action"))
		return
	}

	// Формируем ответ
	ans, err := json.Marshal(JsonResponse{
		Result:  resultAction,
		TimeNow: time.Now(),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("error during marshal: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(ans); err != nil {
		fmt.Println("error writing response: ", err)
	}
}

func main() {
	fmt.Println("Server starting on :8080...")
	http.HandleFunc("/calculate", calculateHande)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("error starting server: ", err)
	}
}
