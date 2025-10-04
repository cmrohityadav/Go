package main

import (
	"fmt"
	"sync"
)


func main(){

// myChan:=make(chan int);
myChan:=make(chan int,2);

wg:=&sync.WaitGroup{}


wg.Add(2);
// R Only (ch <-chan int)
go func(ch chan int,wg *sync.WaitGroup){

	fmt.Println(<-ch);
	
	val,isChanOpen:=<-ch;

	fmt.Println("Val: ",val);
	fmt.Println("isopen: ",isChanOpen);

	wg.Done();
}(myChan,wg);

// send ONLY (ch chan<- int)
go func(ch chan int,wg *sync.WaitGroup){
	close(ch)
	// ch<-5
	// ch<-10

	wg.Done();
}(myChan,wg);

wg.Wait();




}