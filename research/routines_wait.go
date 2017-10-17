package main

import (
	"fmt"
	"sync"
	"time"
)
import "math/rand"

var wg = sync.WaitGroup{}
//var counter = 0

func main(){

fmt.Print(rand.Intn(100))
	
	for i := 0; i < 12; i++ {
		wg.Add(1)
		go sayHello(i)
	}

	wg.Wait()

}

func sayHello(num1 int){
	time.Sleep(1000 * time.Millisecond)
	fmt.Printf("Hello %v\n", num1)
	wg.Done()
}