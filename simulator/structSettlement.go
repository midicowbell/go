package simulator

import (
	"fmt"
	"math/rand"
)

type Settlement struct {
	Name           string
	Resources      int64
	Population     int64
	Buildings      []Building
	BuiltBuildings map[string]bool
}

func (c Settlement) StatusReport() {
	fmt.Printf("---Поселение %s\n", c.Name)
	fmt.Printf("---Количество ресурсов в данном поселении: %d\n", c.Resources)
	fmt.Printf("---Население %d человек\n", c.Population)

}
func (s *Settlement) MakingProfit() {
	if len(s.Buildings) == 0 {
		fmt.Println("У вас нет построек, на этом ходу вы ничего не получаете")
	} else {
		for _, value := range s.Buildings {
			value.CollectIncome(s)
		}
	}

}

func (c *Settlement) ApplyEvent(sub int64) {
	if sub == 0 {
		died := rand.Int63n(10)
		fmt.Printf("В поселении засуха! Умерло %d человек\n", died)
		c.Population -= died
	} else if sub == 1 {
		new1 := rand.Int63n(30)
		fmt.Printf("В поселение прибыли беженцы в количестве %d человек!\n", new1)
		c.Population += new1
	} else if sub == 2 {
		hmnaid := rand.Int63n(25)
		fmt.Printf("В поселение прибыла гуманитарная помощь 🪂 в размере %d ресурсов\n", hmnaid)
		c.Resources += hmnaid
	} else {
		fmt.Println("Сегодня ничего не произошло")
	}
}

type Building interface {
	CollectIncome(s *Settlement)
}
type BuildingStats struct {
	Price int
	Bonus int
	Type  string
}

type Farm struct {
	Price int64
	Bonus int64
	Type  string
}

func NewFarm() Farm {
	return Farm{
		Price: 50,
		Bonus: 8,
		Type:  "gold",
	}
}

func (f Farm) CollectIncome(s *Settlement) {
	fmt.Printf("Ваша ферма принесла %d золота\n", f.Bonus)
	s.Resources += f.Bonus
}

type Mine struct {
	Price int64
	Bonus int64
	Type  string
}

func (m Mine) CollectIncome(s *Settlement) {
	fmt.Printf("Ваша шахта принесла %d золота\n", m.Bonus)
	s.Resources += m.Bonus
}

func NewMine() Mine {
	return Mine{
		Price: 25,
		Bonus: 4,
		Type:  "gold",
	}
}

type House struct {
	Price int64
	Bonus int64
	Type  string
}

func NewHouse() House {
	return House{
		Price: 15,
		Bonus: 3,
		Type:  "people",
	}
}

func (h House) CollectIncome(s *Settlement) {
	fmt.Printf("Ваш домик принес %d жителей\n", h.Bonus)
	s.Population += h.Bonus
}
