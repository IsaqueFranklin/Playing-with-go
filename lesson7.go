//CHANNELS IN GO: WHAT ARE AND HOW THEY WORK => They are goroutines to pass around infomartion.
//THey hold data (an integer, a slice, anything)
//They're thread safe, avoid data races when reading memory.
//We can listen when data is added or removed from a channel and block code executions until one the these events arrives.

package main
import (
	"fmt"
	"time"
)

func main(){
	var c = make(chan int, 5) //This is a buffer channel
	go process(c)
	for i:= range c{
		fmt.Println(i)
		time.Sleep(time.Second*1)
	}
}

func process(c chan int){
	defer 	close(c)
	for i:=0; i<5; i++{
		c <- i
	}
	fmt.Println("Exiting process")
}