package main

import (
	"context"
	"fmt"
	"math/rand"
	"study/logs"
	"sync"
	"time"
)

func main() {
	mtx := sync.RWMutex{}
	wgSender := sync.WaitGroup{}
	wgWorker := sync.WaitGroup{}
	ctxGenerator, ctxCancel := context.WithCancel(context.Background())
	logStats := make(map[int]string)
	logChan := make(chan logs.LogEntry)
	for i := 0; i < 10; i++ {
		wgSender.Add(1)
		go func() {
			defer wgSender.Done()
			logs.Generator(ctxGenerator, rand.Int(), logChan)
		}()
	}
	go func() {
		wgSender.Wait()
		close(logChan)
	}()
	for i := 0; i < 5; i++ {
		wgWorker.Add(1)
		go logs.Worker(i, logChan, logStats, &mtx, &wgWorker)
	}
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		mtx.RLock()
		fmt.Printf("Текущее состояние лога на %d секунду: %v\n", i, logStats)
		mtx.RUnlock()
	}
	ctxCancel()
	wgWorker.Wait()

	fmt.Println("Итоговые логи:")
	for key, value := range logStats {
		fmt.Printf("Service %d: %s\n", key, value)
	}
}
