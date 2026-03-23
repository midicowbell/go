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
		Type:  "gold",
	},
	"Охотничий домик": {
		Price: 20,
		Bonus: 5,
		Type:  "gold",
	},
	"Деревянный домик": {
		Price: 15,
		Bonus: 5,
		Type:  "people",
	},
	"Кузница": {
		Price: 25,
		Bonus: 10,
		Type:  "gold",
	},
}

func MakingProfit(b map[string]bool, mySettlement *Settlement) {
	for name, isBuilt := range b {
		if isBuilt {
			effect := buildingEffects[name]
			fmt.Printf("%s принес бонус: %d\n", name, effect.Bonus)

			switch effect.Type {
			case "gold":
				mySettlement.Resources += int64(effect.Bonus)
			case "people":
				mySettlement.Population += int64(effect.Bonus)
			}
		}
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
		if mySettlement.Population <= 0 {
			fmt.Println("Игра завершена! Деревная вымерла")
			return
		}
		text := scanner.Text()
		fields := strings.Fields(text)
		if len(fields) != 0 {
			if fields[0] == "статус" {
				mySettlement.StatusReport()
			} else if fields[0] == "строить" && len(fields) >= 2 {
				buildingName := strings.Join(fields[1:], " ")
				stats, exists := buildingEffects[buildingName]
				if !exists {
					fmt.Println("Такого здания не сущесвует!")
					continue
				}
				if buildings[buildingName] {
					fmt.Println("Это здание уже есть в вашем поселении!")
					continue
				}
				if mySettlement.Resources < int64(stats.Price) {
					fmt.Printf("Недостаточно золота! Нужно %d, а у вас %d\n", stats.Price, mySettlement.Resources)
					continue
				}
				buildings[buildingName] = true
				mySettlement.Resources -= int64(stats.Price)
				fmt.Printf("Вы успешно построили: %s\n", buildingName)
				history = append(history, "Построено: "+buildingName)
			} else if fields[0] == "ход" {
				event := rand.Int63n(5)
				mySettlement.ApplyEvent(event)
				MakingProfit(buildings, &mySettlement)
			} else if fields[0] == "история" {
				if len(history) == 0 {
					fmt.Println("История пуста")
				} else {
					for _, value := range history {
						fmt.Println(value)
					}
				}

			} else {
				fmt.Println("Команда неизвестна")
			}
		}
	}
}
