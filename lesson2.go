package main
import (
	"errors"	
	"fmt"
)

func main(){
	printMe("Stack sats!")

	var numerator float64 = 11568.56897
	var denominator float64 = 6.896

	var result, multiplication, err = intDivision(numerator, denominator)
	
	switch { 
		case err!=nil:
			fmt.Printf(err.Error())
		default:
			fmt.Printf("The result of float64 division is %v and the multiplication of float64 is %v.", result, multiplication)
	}

	switch multiplication {
		case 0:
			fmt.Printf("Houve multiplicação por zero.")
		case 59:
			fmt.Printf("HAHAHAHAHA")
		default:
			fmt.Printf("Deu bom.")
	}
	
	if err!=nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Printf("The result of float64 division is %v and the multiplication of float64 is %v.", result, multiplication)
	}
}

func printMe(printValue string){
	fmt.Println(printValue)
}

func intDivision(numerator float64, denominator float64) (float64, float64, error) {
	var err error //Default type is nil

	var multiplication float64 = numerator*denominator

	if denominator==0 {
		err = errors.New("Cannot divide by zero.")

		return 0, multiplication, err

	} else {
		var result float64 = numerator/denominator

		return result, multiplication, err
	}
}