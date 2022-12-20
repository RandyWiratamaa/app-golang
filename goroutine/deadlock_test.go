package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type UserBelance struct {
	sync.Mutex
	Name    string
	Balance int
}

// Method untuk Lock(menggunakan sync.Mutex)
func (user *UserBelance) Lock() {
	user.Mutex.Lock()
}

// Method untuk Unlock(menggunakan sync.Mutex)
func (user *UserBelance) Unlock() {
	user.Mutex.Unlock()
}

// Method untuk change pada balance
func (user *UserBelance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBelance, user2 *UserBelance, amount int) {
	user1.Lock()
	fmt.Println("User 1 Lock", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("User 2 Lock", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := UserBelance{
		Name:    "Randy",
		Balance: 2000000,
	}

	user2 := UserBelance{
		Name:    "Wiratama",
		Balance: 1000000,
	}

	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 500000)

	time.Sleep(3 * time.Second)

	fmt.Println("User ", user1.Name, ", Balance ", user1.Balance)
	fmt.Println("User ", user2.Name, ", Balance ", user2.Balance)
}

func TestSyncPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}

	pool.Put("Randy")
	pool.Put("Wiratama")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}
	time.Sleep(11 * time.Second)
	fmt.Println("Selesai")
}

func addToMap(data *sync.Map, nilai int, group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	data.Store(nilai, nilai)
}

func TestSyncMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go addToMap(data, i, group)
	}

	group.Wait()
	data.Range(func(key, nilai interface{}) bool {
		fmt.Println(key, ":", nilai)
		return true
	})
}
