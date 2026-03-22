package simulator

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var buildingEffects = map[string]BuildingStats{
	"Ферма": {
		Price: 50,
		Bonus: 10,
	},
	"Охотничий домик": {
		Price: 20,
		Bonus: 5,
	},
	"Деревянный домик": {
		Price: 15,
		Bonus: 5,
	},
	"Кузница": {
		Price: 25,
		Bonus: 10,
	},
}

func MakingProfit(b map[string]bool, mySettlement *Settlement) {
	if b["Ферма"] == true {
		fmt.Printf("Ферма принесла %d ресурсов\n", buildingEffects["Ферма"].Bonus)
		mySettlement.Resources += int64(buildingEffects["Ферма"].Bonus)
	}
	if b["Охотничий домик"] == true {
		fmt.Printf("Охотничий домик принес %d золота\n", buildingEffects["Охотничий домик"].Bonus)
		mySettlement.Resources += int64(buildingEffects["Охотничий домик"].Bonus)
	}
	if b["Деревянный домик"] == true {
		fmt.Printf("Деревянный домик принес %d людей\n", buildingEffects["Деревянный домик"].Bonus)
		mySettlement.Population += int64(buildingEffects["Деревянный домик"].Bonus)
	}
	if b["Кузница"] == true {
		fmt.Printf("Кузница принесла %d золота\n", buildingEffects["Кузница"].Bonus)
		mySettlement.Resources += int64(buildingEffects["Кузница"].Bonus)
	}
}

// Ферма + 10 к ресурсам, охотничий домик + 5 к ресурсам, деревянный дом - 10 к ресурасам
// Кузница + 10 к ресурсам

func Play() {
	fmt.Println("------Симулятор поселения------")
	fmt.Println("Список доступных комманд: сторить [здание], статус, история, ход")
	fmt.Println("С помощью команды сторить можно возвести новое здание")
	fmt.Println("С помощью команды статус можно вывести статус вашего поселения")
	fmt.Println("С помощью команды история можно вывести список ваших изменений")
	fmt.Println("С помощью команды ход можно перейти в другой день")
	mySettlement := Settlement{Name: "Моя Деревня", Resources: 100, Population: 10, Status: true}
	history := make([]string, 0, 10)
	buildings := map[string]bool{"Ферма": false, "Охотничий домик": false, "Деревянный домик": false, "Кузница": false}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			fmt.Println("Возникла ошибка при чтении твоей команды, попробуй еще раз")
		}
		text := scanner.Text()
		fields := strings.Fields(text)
		if len(fields) != 0 {
			if fields[0] == "статус" {
				mySettlement.StatusReport()
			} else if fields[0] == "строить" && len(fields) >= 2 {
				buildingName := strings.Join(fields[1:], " ")
				if buildingName == "Ферма" {
					val := buildings["Ферма"]
					if !val {
						buildings["Ферма"] = true
						fmt.Printf("Вы построили ферму за %d золота\n", buildingEffects["Ферма"].Price)
						mySettlement.Resources -= int64(buildingEffects["Ферма"].Price)
						history = append(history, "Вы построили ферму")
					} else {
						fmt.Println("Ферма уже построена!!!")
					}
				} else if buildingName == "Охотничий домик" {
					val := buildings["Охотничий домик"]
					if !val {
						buildings["Охотничий домик"] = true
						fmt.Printf("Вы построили охотничий домик за %d золота\n", buildingEffects["Охотничий домик"].Price)
						mySettlement.Resources -= int64(buildingEffects["Охотничий домик"].Price)
						history = append(history, "Вы построили охотничий домик")
					} else {

					}
				} else if buildingName == "Деревянный домик" {
					val := buildings["Деревянный домик"]
					if !val {
						buildings["Деревянный домик"] = true
						fmt.Printf("Вы построили деревянный домик за %d золота\n", buildingEffects["Деревянный домик"].Price)
						mySettlement.Resources -= int64(buildingEffects["Деревянный домик"].Price)
						history = append(history, "Вы построили деревянный домик")
					} else {

					}
				} else if buildingName == "Кузница" {
					val := buildings["Кузница"]
					if !val {
						buildings["Кузница"] = true
						fmt.Printf("Вы построили кузницу за %d золота\n", buildingEffects["Кузница"].Price)
						mySettlement.Resources -= int64(buildingEffects["Кузница"].Price)
						history = append(history, "Вы построили кузницу")
					} else {

					}
				} else {
					fmt.Println("Такого здания нет в списке")
				}
			} else if fields[0] == "ход" {
				event := rand.Int63n(5)
				mySettlement.ApplyEvent(event)
				MakingProfit(buildings, &mySettlement)
			} else if fields[0] == "история" {
				for _, value := range history {
					fmt.Println(value)
				}
			}
		}
	}
}
