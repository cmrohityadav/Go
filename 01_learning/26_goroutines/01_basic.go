package main

import (
	"fmt"
	"sync"
	"time"
)
func task(id int){
	fmt.Println("Doing task ",id);
}

var signals=[]string{"test"}



func main(){

	for i:=0;i<=10;i++ {
		go task(i)
	}

	time.Sleep(time.Second*2)

// ---------02 code----------------

	go Greeter("hello");
	Greeter("world");

//--------03 code (Wait group) ----------
websitesList:=[]string{"https://github.com/cmrohityadav","https://google.com","https://facebook.com",}

var wg sync.WaitGroup //In real project we create pointer of it


for _,web:= range websitesList{
	wg.Add(1)
	go GetStatusCode(web,&wg);
}
wg.Wait();

fmt.Println("signals :",signals);



//--------Mutex-----------

var coruptCounter int=0;

for i := 0; i < 5000; i++ {
    go func() {
       coruptCounter++;
    }()
    go func() {
        coruptCounter++;
    }()
}

time.Sleep(2 * time.Second)

fmt.Println("corupt Counter:",coruptCounter);

var mut sync.Mutex //In real project we create pointer of it

var safeCounter int = 0

for i := 0; i < 5000; i++ {
    go func() {
        mut.Lock()
        safeCounter++
        mut.Unlock()
    }()
    go func() {
        mut.Lock()
        safeCounter++
        mut.Unlock()
    }()
}

time.Sleep(2 * time.Second)
fmt.Println("safe Counter:", safeCounter)

}
