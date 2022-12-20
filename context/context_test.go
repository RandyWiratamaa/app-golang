package context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	contextBackground := context.Background()
	fmt.Println("Context Background :", contextBackground)
	contextTodo := context.TODO()
	fmt.Println("Context TODO :", contextTodo)
}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "Parent A-Context B")
	contextC := context.WithValue(contextA, "c", "Parent A-Context C")

	contextD := context.WithValue(contextB, "d", "Parent B-Context D")
	contextE := context.WithValue(contextB, "e", "Parent B-Context E")

	contextF := context.WithValue(contextC, "f", "Parent C-Context F")

	fmt.Println("Context A :", contextA)
	fmt.Println("Context B :", contextB)
	fmt.Println("Context C :", contextC)
	fmt.Println("Context D :", contextD)
	fmt.Println("Context E :", contextE)
	fmt.Println("Context F :", contextF)

	fmt.Println("1.", contextF.Value("f"))
	fmt.Println("2.", contextE.Value("b"))
	fmt.Println("3.", contextE.Value("d")) // sesama child diparent yang sama tidak bisa saling mengambil data
	fmt.Println("4.", contextE.Value("c")) // Context C bukan parent dari Context E (Beda Parent)
	fmt.Println("5.", contextA.Value("b")) // Context A tidak bisa mengambil data child
	fmt.Println("6.", contextB.Value("d")) // Context B tidak bisa mengambil data child
}

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second) // Simulate Slow Process
			}
		}
	}()
	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Total Goroutine Awal : ", runtime.NumGoroutine())
	parentCtx := context.Background()
	ctx, cancel := context.WithCancel(parentCtx)

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("Counter : ", n)
		if n == 10 {
			break
		}
	}
	cancel()

	fmt.Println("Total Goroutine ketika berhenti : ", runtime.NumGoroutine())
}

func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Total Goroutine Awal : ", runtime.NumGoroutine())
	parentCtx := context.Background()
	ctx, cancel := context.WithTimeout(parentCtx, 5*time.Second)
	defer cancel()

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("Counter : ", n)
	}
	fmt.Println("Timeout! Goroutine telah berhenti...")
	fmt.Println("Total Goroutine ketika berhenti : ", runtime.NumGoroutine())
}

func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Total Goroutine Awal : ", runtime.NumGoroutine())
	parentCtx := context.Background()
	ctx, cancel := context.WithDeadline(parentCtx, time.Now().Add(5*time.Second))
	defer cancel()

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("Counter : ", n)
	}
	fmt.Println("Timeout! Goroutine telah berhenti...")
	fmt.Println("Total Goroutine ketika berhenti : ", runtime.NumGoroutine())
}
