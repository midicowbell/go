package logs

import (
	"context"
	"fmt"
	"math"
	"sync"
	"time"
)

var logsErros [3]string = [3]string{"INFO", "WARNING", "ERROR"}

type LogEntry struct {
	serviceID int
	level     string
	message   string
}

func Generator(ctx context.Context, id int, logChan chan<- LogEntry) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Генератор с ID %d завершает работу!\n", id)
			return
		default:
			logChan <- LogEntry{
				serviceID: id,
				level:     logsErros[id%3],
				message:   "Детекция предупреждения",
			}
			time.Sleep(500 * time.Millisecond)
		}
	}

}
func Worker(id int, logChan <-chan LogEntry, stats map[int]string, mtx *sync.RWMutex, wg *sync.WaitGroup) {
	defer wg.Done()
	for log := range logChan {
		mtx.Lock()
		stats[log.serviceID] = log.level
		mtx.Unlock()
		time.Sleep(500 * time.Millisecond)
	}

}

func IsPrime(n int) bool {
	for i := 1; i < int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
