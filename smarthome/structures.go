package smarthome

import (
	"fmt"
	"time"
)

type Device interface {
	TurnOn()
	TurnOff()
}

func ManageDevice(d Device) {
	d.TurnOn()
	d.TurnOff()
}

type Lamp struct{}

func (l Lamp) TurnOn() {
	fmt.Println("Включаю лампочку")
	for i := 0; i < 5; i++ {
		fmt.Print(".")
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println()
	fmt.Println("Лампочка включена")
}
func (l Lamp) TurnOff() {
	fmt.Println("Выключаю лампочку")
	for i := 0; i < 5; i++ {
		fmt.Print(".")
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println()
	fmt.Println("Лампочка выключена")
}

type AirConditioner struct{}

func (airCond AirConditioner) TurnOn() {
	fmt.Println("Стало жарко? Включаю кондиционер")
	for i := 0; i < 5; i++ {
		fmt.Print(".")
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println()
	fmt.Println("Кондиционер включен")
}
func (airCond AirConditioner) TurnOff() {
	fmt.Println("Стало холодно? Выключаю кондиционер")
	for i := 0; i < 5; i++ {
		fmt.Print(".")
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println()
	fmt.Println("Кондиционер выключен")
}
