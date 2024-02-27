package main
import "fmt"

type gasEngine struct{
	mpg uint8
	gallons uint8
}

func main(){
	var myEngine gasEngine
	fmt.Println(myEngine.mpg, myEngine.gallons)
}