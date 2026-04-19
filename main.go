package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"
)

var totalCount atomic.Int64

func handlerReset(w http.ResponseWriter, r *http.Request) {
	totalCount.Add(0)
	msg := "Число было сброшено. Последнее значение: " + strconv.Itoa(int(totalCount.Load()))
	_, err := w.Write([]byte(msg))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("fail to write HTTP response: ", err)
	}
}

func handlerAdd(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		msg := "fail to read HTTP body" + err.Error()
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to wrtie HTTP response: ", err)
		}
		return
	}
	httpRequestBodyString := string(httpRequestBody)
	add, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := "fail to convert HTTP body to int" + err.Error()
		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("fail to wrtie HTTP response: ", err)
		}
		return
	}
	totalCount.Add(int64(add))
	msg := "К числу было добавлено: " + strconv.Itoa(add)
	fmt.Println(msg)
	_, err = w.Write([]byte(msg))
	if err != nil {
		fmt.Println("fail to write HTTP response: ", err)
	}
}

func handlerGet(w http.ResponseWriter, r *http.Request) {
	ans := int(totalCount.Load())
	msg := "Текущее число = " + strconv.Itoa(ans)
	_, err := w.Write([]byte(msg))
	if err != nil {
		fmt.Println("fail to wrtie HTTP response: ", err)
		return
	}
}

func main() {
	totalCount.Add(0)
	http.HandleFunc("/add", handlerAdd)
	http.HandleFunc("/reset", handlerReset)
	http.HandleFunc("/value", handlerGet)
	fmt.Print("Запускаю HTTP сервер")
	for i := 0; i < 5; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Print(".")
		if i == 4 {
			fmt.Println()
		}
	}
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Завершаю программу")
}
