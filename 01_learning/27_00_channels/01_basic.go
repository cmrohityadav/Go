package main

import (
    "fmt"
    "math/rand"
    "time"
)

func processNum(numChan chan int) {
    fmt.Println("processing number", <-numChan)
}

func randomNumPrint(Channel chan int) {
    for val := range Channel {
        fmt.Println("Value:", val)
		time.Sleep(1*time.Second)
    }
	

}


func sum(result chan int,num1 ,num2 int){
    sum:=num1+num2;

    result<-sum;
}

func main() {
    
    rand.Seed(time.Now().UnixNano())

    numChan := make(chan int)
    go processNum(numChan)

    time.Sleep(5 * time.Second)
    numChan <- 5
    time.Sleep(1 * time.Second)


    randNumChan := make(chan int)
	
    go randomNumPrint(randNumChan)

    sumChannel:=make(chan int)

    go sum(sumChannel,10,5);

    fmt.Println("From Sum channel: ",<-sumChannel)


    for {
        randNumChan <- rand.Intn(100)
    }
}
