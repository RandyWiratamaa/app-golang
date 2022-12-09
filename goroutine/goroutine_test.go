package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Golang Goroutine")
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestCreateGoroutine(t *testing.T) {
	// goroutine diawali dengan go
	go RunHelloWorld()
	fmt.Println("ups")

	time.Sleep(1 * time.Second)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		// goroutine diawali dengan go
		go DisplayNumber(i)
	}
	time.Sleep(10 * time.Second)
}
