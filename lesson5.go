//POINTERS AND WHAT ARE POINTERS

package main
import "fmt"

func main(){
	//My pointer cannot be nil for me to assing values to it, therefore I'm initializing it here.
	var p *int32 = new(int32)
	var i int32
	fmt.Printf("The value p points to is: %v", *p)
	fmt.Printf("\nThe value if i is: %v", i)
	//*p = 10 -> Colocando um valor no ponteiro p
	p = &i //ponteiro p agora está apondando para o valor de i

	//Se eu modifico o valor do ponteiro *p, o valor de i também muda
	*p=1
	fmt.Printf("\nThe value p points to is: %v", *p)
	fmt.Printf("\nThe value if i is: %v", i)
}