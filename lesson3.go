package main
import "fmt"

func main(){
	//Arrays are fixed length, same type, indexable, contiguous in memory.
	var intArr [3]int32
	fmt.Println(intArr[0])
	

	var intSlice []int32 = []int32{4,5,6}
	fmt.Println(intSlice)
	intSlice = append(intSlice, 7)

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam":23, "Sarah":45, "Bruce": 32}
	fmt.Println(myMap2["Adam"])
	var age, ok = myMap2["Jason"]
	delete(myMap2, "Adam")
	if ok{
		fmt.Printf("The age is %v.", age)
	} else {
		fmt.Println("invalid Name.")
	}

	for name, age:= range myMap2{
		fmt.Printf("Name: %v, Age: %v \n", name, age)
	}
}