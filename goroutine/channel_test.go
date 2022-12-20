package goroutine

import (
	"fmt"
	"strconv"
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

	// mengambil data dari channel
	data := <-channel // string Randy Wiratama dimasukkan ke dalam variable data
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

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	//Mengirim data atau channel
	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	// Range pada saat menerima data
	for data := range channel {
		fmt.Println("Menerima data", data)
	}

	fmt.Println("Selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel ke 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel ke 2", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}
}

func TestDefaultChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel ke 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel ke 2", data)
			counter++
		default:
			fmt.Println("Menunggu Data...")
		}
		if counter == 2 {
			break
		}
	}
}
