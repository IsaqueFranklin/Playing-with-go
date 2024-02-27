package main
import "fmt"
import "unicode/utf8"

func main(){
	var intNum int = 32767
	//I can choose how much store I use to store my int number (int8, int16, int32, int64 and so on)
	//I cant also use uint that it a int that only stores positive integers.
	fmt.Println(intNum)

	//I need to specify the float as float32 or float64, there is not only "float".
	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)

	var floatNum32 float32 =  10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var myString string = "Hello" + " " + "world!"
	fmt.Println(myString)

	fmt.Println(utf8.RuneCountInString("Y"))

	var myBoolean bool = false
	fmt.Println(myBoolean)
	myBoolean = true
	fmt.Println(myBoolean)

	//The default value for all the ints, unsigned ints, floats and rune is 0, for strings is "" and bool is false.

	//When the string content is declared right way there is no need to specify string.

	myVar := "text" // abreviação de var myVar = "text" 
	fmt.Println(myVar)

	var1, var2 := 1, 2 //var var1, var int = 1, 2 //Iniciando multiplas variáveis de uma vez.
	fmt.Println(var1,var2) 

	//const is also usable but its value cannot change down the road, its fixed.
}