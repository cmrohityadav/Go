package main

import (
	"fmt"
	"net/http"
	"sync"
)

func GetStatusCode(endpoint string,wg *sync.WaitGroup){
	defer wg.Done()

	res,err:=http.Get(endpoint);
	if err!= nil{
		fmt.Println("Oops in endpoint");
	}else{
		signals=append(signals, endpoint);
	fmt.Println("Status code for website: ",endpoint," : ",res.StatusCode);

	}


}
