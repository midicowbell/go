package simulator

import (
	"fmt"
	"math/rand"
)

type Settlement struct {
	Name       string
	Resources  int64
	Population int64
	Status     bool
}

func (c Settlement) StatusReport() {
	fmt.Printf("---Поселение %s\n", c.Name)
	fmt.Printf("---Количество ресурсов в данном поселении: %d\n", c.Resources)
	fmt.Printf("---Население %d человек\n", c.Population)

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
	} else {
		fmt.Println("Сегодня ничего не произошло")
	}
}

type BuildingStats struct {
	Price int
	Bonus int
	Type  string
}
