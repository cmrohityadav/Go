package main

import (
	"fmt"
	"sync"
	"time"
)

func goThreadFunction(name int, wg *sync.WaitGroup) {
	defer wg.Done();
	fmt.Println("started thread: ", name);
	time.Sleep(time.Second*2);
	fmt.Println("end thread: ",name);
}


func main() {
	var wg sync.WaitGroup;

	for i:= range 20{
		wg.Add(1);
		go goThreadFunction(i,&wg)
	}

	wg.Wait()
}
