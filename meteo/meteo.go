package meteo

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type WeatherReport struct {
	StationID int32
	Temp      int
}

func sensors(ctx context.Context, wg *sync.WaitGroup, id int32, dataChan chan<- WeatherReport) {
	defer wg.Done()
	select {
	case <-ctx.Done():
		fmt.Println("Замер данных прекращен")
	case <-time.After(1000 * time.Millisecond):
		fmt.Println("Сенсор с id: ", id, "начал работу")
		dataChan <- WeatherReport{
			StationID: id,
			Temp:      rand.Intn(30),
		}
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Сенсор с id: ", id, "завершил работу")
	}

}

func SensorsPool(ctx context.Context, id int32) <-chan WeatherReport {
	dataChan := make(chan WeatherReport)
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go sensors(ctx, &wg, rand.Int31(), dataChan)
	}
	go func() {
		wg.Wait()
		close(dataChan)
	}()
	return dataChan
}
