package goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}

	for i := 1; i <= 1000; i++ {
		go func() {
			group.Add(1)
			for j := 1; j <= 100; j++ {
				atomic.AddInt64(&x, 1)
			}
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter : ", x)
}

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println("Sebelum : ", time.Now())

	time := <-timer.C
	fmt.Println("Sesudah : ", time)
}

func TestTimerAfter(t *testing.T) {
	channel := time.After(5 * time.Second)
	fmt.Println("Sebelum : ", time.Now())

	time := <-channel
	fmt.Println("Sesudah : ", time)
}

func TestTimerAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}

	group.Add(1)
	time.AfterFunc(5*time.Second, func() {
		fmt.Println("Sesudah : ", time.Now())
		group.Done()
	})
	fmt.Println("Sebelum : ", time.Now())
	group.Wait()
}

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(2 * time.Second)

	go func() {
		time.Sleep(10 * time.Second)
		ticker.Stop()
	}()

	for tick := range ticker.C {
		fmt.Println(tick)
	}
}

func TestGomaxprocs(t *testing.T) {
	totalCpu := runtime.NumCPU()
	runtime.GOMAXPROCS(20) // Mengubah total Thread
	totalThread := runtime.GOMAXPROCS(-1)
	totalGoroutine := runtime.NumGoroutine()

	fmt.Println("Total CPU : ", totalCpu)
	fmt.Println("Total Thread : ", totalThread)
	fmt.Println("Total Goroutine : ", totalGoroutine)
}
