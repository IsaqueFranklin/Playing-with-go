package main
import "fmt"

func main(){
	var slice = []int32{1,2,3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Println(slice)
	fmt.Println(sliceCopy)

	//Under the hood slices contain pointers to an underlying array.
}