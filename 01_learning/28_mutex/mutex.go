package main

import (
	"fmt"
	"sync"
)

type post struct{

	views int
	mu sync.Mutex
}

func(p *post) inc(wg *sync.WaitGroup){
	defer func ()  {
			wg.Done()
			p.mu.Unlock()	

	}()
	

	p.mu.Lock()
	p.views+=1;
	// p.mu.Unlock()
}

func main(){

	myPost:=post{views: 0};
	var wg sync.WaitGroup
	for i:=0;i<10000;i++ {
		wg.Add(1)
		go myPost.inc(&wg)
	}

	wg.Wait()

	fmt.Println(myPost.views)

}