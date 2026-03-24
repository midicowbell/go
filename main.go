package main

import (
	"fmt"
	"study/simulator"
)

func main() {
	err := simulator.Play()
	if err != nil {
		fmt.Printf("Ошибка: %s", err.Error())
	}
}
