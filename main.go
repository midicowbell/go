package main

import (
	"messengers/messenger"
)

func main() {
	t1 := messenger.Telegram{
		Username: "Jack",
	}
	w1 := messenger.WhatsApp{}
	slice := make([]messenger.Messenger, 0, 2)
	slice = append(slice, t1)
	slice = append(slice, w1)

	for _, value := range slice {
		value.SendMessage("Привет")
	}
}
