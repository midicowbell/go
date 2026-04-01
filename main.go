package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"time"
)

func SpamBot(ctx context.Context, id int32) {
	for {
		countLetters := rand.IntN(20)
		select {
		case <-ctx.Done():
			fmt.Printf("Спам-бот с ID [%d] прекратил свою работу\n", id)
			return
		default:
			fmt.Printf("Спам-бот с ID [%d] отправил %d письмем\n", id, countLetters)
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	parentContext, parentCancel := context.WithCancel(context.Background())
	for i := 0; i < 3; i++ {
		genId := rand.Int32()
		go SpamBot(parentContext, genId)
	}

	time.Sleep(2 * time.Second)

	parentCancel()

	time.Sleep(200 * time.Millisecond)
}
