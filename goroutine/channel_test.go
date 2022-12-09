package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func GiveMeResponse(var_channel chan string) {
	time.Sleep(2 * time.Second)

	var_channel <- "Randy Wiratama"
}

func ChannelOnlyIn(var_channel chan<- string) {
	time.Sleep(2 * time.Second)

	var_channel <- "Randy"
}

func ChannelOnlyOut(var_channel <-chan string) {
	data := var_channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	var_channel := make(chan string)
	defer close(var_channel)

	go ChannelOnlyIn(var_channel)
	go ChannelOnlyOut(var_channel)

	time.Sleep(3 * time.Second)
	close(var_channel)
}

func TestChannelAsParameter(t *testing.T) {
	//deklarasi variable channel
	// type data channel adalah chan
	var_channel := make(chan string)
	defer close(var_channel)
	//memanggil function
	go GiveMeResponse(var_channel)

	// string Randy Wiratama dimasukkan ke dalam variable data
	// mengambil data dari channel
	data := <-var_channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
	close(var_channel)
}

func TestMakeChannel(t *testing.T) {
	//deklarasi variable channel
	// type data channel adalah chan
	channel := make(chan string)
	defer close(channel)
	//anonymous function
	go func() {
		time.Sleep(2 * time.Second)
		// mengirim ke channel
		channel <- "Randy Wiratama"
	}()

	// string Randy Wiratama dimasukkan ke dalam variable data
	// mengambil data dari channel
	data := <-channel
	fmt.Println(data)
	close(channel)
}

func TestBufferedChannel(t *testing.T) {
	//deklarasi variable
	channel := make(chan string, 3)
	defer close(channel)
	// go ChannelOnlyIn(channel)
	channel <- "Randy"
	channel <- "Wiratama"

	fmt.Println("Selesai")
}
