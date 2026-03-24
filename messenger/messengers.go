package messenger

import "fmt"

type Messenger interface {
	SendMessage(text string)
}

func Send(m Messenger, text string) {
	m.SendMessage(text)
}

type Telegram struct {
	Username string
}

func (t Telegram) SendMessage(text string) {
	fmt.Println("Тебе пришло новое сообщение от ", t.Username)
	fmt.Println(text)
}

type WhatsApp struct {
}

func (w WhatsApp) SendMessage(text string) {
	fmt.Println(text)
}
