package main

import (
	"fmt"
	"time"
	"sync"
)

var wg = sync.WaitGroup{}

func main(){

	var msg = "hello"

	wg.Add(1)
	
	go func(msg string){
		fmt.Println(msg)
		wg.Done()
	}(msg)


	msg = "Goodbye"

	wg.Wait()

	time.Sleep(100 * time.Millisecond)
}
