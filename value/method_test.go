package value

import (
	"fmt"
	"testing"
	"time"
)

func TestMethod(t *testing.T) {
	slice1 := []string{"zhang", "san"}
	modify(slice1)
	fmt.Println(slice1)
}

func modify(data []string) {
	data = nil
}


func TestChannel(t *testing.T) {
	c := make(chan bool, 10)
	fmt.Println("channel address", &c)
	for i := 0; i < 10; i++ {
		go Go(c, i)
	}
	time.Sleep(20*time.Second)
}
func Go(c chan bool, index int) {
	sum := 0
	for i := 0; i < 1000000; i++ {
		sum += i
	}
	fmt.Println(sum)
	fmt.Println(c)
	c <- true
}

func TestBufferChannel(t *testing.T) {

	message := make(chan string, 2)

	message <- "abd"
	message <- "dba"


	fmt.Print(<-message)
	fmt.Print(<-message)


	//if use range to get the value in channel
	//there gonna be deadlock error

	//for value := range message{
	//	fmt.Println(value)
	//}

}