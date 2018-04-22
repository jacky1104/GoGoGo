package channel

import (
	"fmt"
	"testing"
	"time"
)

func sendToChannel(a chan int, value int) {

	for i := 0; i < value; i++ {
		a <- i
	}

	close(a)
}


func readFromChannel(a chan int, sync chan bool) {

	for key := range a{
		fmt.Println("From channel:", key)
	}

	sync <- true
}

func TestChannel(t *testing.T) {

	channelOne := make(chan int, 10)
	sync := make(chan bool)

	go sendToChannel(channelOne,7)

	go readFromChannel(channelOne,sync)

	<-sync

	fmt.Println("Channel Over")

}




func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x + y
		case <-quit:
			fmt.Println("quit")
			return
		case <- time.After(5 * time.Second):
			println("timeout")
			quit <- 0
		}
	}
}

func TestSelect(t *testing.T) {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		//quit <- 0
	}()
	fibonacci(c, quit)
}

func Test123(t *testing.T){
	//if change 2 to 1, there will be panic
	a := make(chan string,2)

	a <- "123"

	b := <-a
	fmt.Printf(b)
}