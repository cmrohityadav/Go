package main

import "net/http"

func ping(w http.ResponseWriter,r *http.Request){
	if r.Method!=http.MethodGet{
		http.Error(w,"Only Get Method is Allow",http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong!!!"))

}

func main() {
	http.HandleFunc("/ping",ping)

	http.ListenAndServe(":8001",nil)
}