package main
import "fmt"

type gasEngine struct{
	mpg uint8
	gallons uint8
	owner //ownerInfo owner is also an option
}

type owner struct {
	name string
}

func main(){
	var myEngine gasEngine = gasEngine{25, 15, owner{"Isaque"}} //If I ommit them, the fields are assigned values in order.
	myEngine.mpg = 20
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.name) //I can also use myEngine.ownerInfo.name
}