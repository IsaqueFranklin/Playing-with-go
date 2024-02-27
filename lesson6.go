//GO ROUTINES is way to launch multiple function and have them execute concurrently.
//Cuncurrency is different from parallel execution.
//Concurrency is multiple tasks in progress at the same time => Using cpu time to move between taks that execute more quickly 
//while waiting for the others to fulfill.

package main
import (
	"fmt"
	//"math/rand"
	"time"
	"sync"
)

var m = sync.RWMutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main(){
	t0 := time.Now()
	for i:=0; i<len(dbData); i++{
		wg.Add(1)
		go dbCall(i) //go na frente mostra para o compilador que quero rodar essa função concorrentemente.
	}
	wg.Wait()
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
	fmt.Printf("\nThe results are %v", results)
}

func dbCall(i int){
	//Simulate DB call delay
	var delay float32 = 2000
	time.Sleep(time.Duration(delay)*time.Millisecond)
	fmt.Println("The result from the database is: ", dbData[i])
	save(dbData[i])
	log()
	wg.Done()
}

func save(result string){
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log(){
	m.RLock()
	fmt.Printf("\nThe current results are: %v", results)
	m.RUnlock()
}