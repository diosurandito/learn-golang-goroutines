package golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Dio Surandito"
		fmt.Println("Selesai mengirim data ke Channel")
	}()

	data := <-channel
	fmt.Println(data)
	// close(channel)
	time.Sleep(5 * time.Second)

}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Dio Surandito"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)
	data := <-channel
	fmt.Println(data)
}
