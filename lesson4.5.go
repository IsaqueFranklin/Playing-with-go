package main
import "fmt"

//Those are structu for a electric and gas car
type gasEngine struct{
	mpg uint8
	gallons uint8
}

type electricEngine struct {
	mpkwh uint8
	kwh uint8
}

//Those are different from normal functions, those are methods that are tied their respectives struct
func (e gasEngine) milesLeft() uint8 {
	return e.gallons*e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh*e.mpkwh
}

//Using interfaces to solve problems
type engine interface {
	milesLeft() uint8
}

//Using the interface the function below can be user for both the kinds of structs, if the struct has a milesLeft() method.
func canMakeIt(e engine, miles uint8){
	if miles<=e.milesLeft(){
		fmt.Printf("You can make it there!")
	} else {
		fmt.Println("Need to fuel up first!")
	}
}

func main(){
	var myEngine electricEngine = electricEngine{25, 15} //If I ommit them, the fields are assigned values in order.
	canMakeIt(myEngine, 50)
}

//Really cool use of golang