package main

import (
	"fmt"
	"net/http"
	"time"
)

var (
	requestCount int
	windowStart = time.Now()
)

func rateLimiter(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// time.Since(windowStart) tab se abi tak ka time difference: time.Now() - start
		if time.Since(windowStart) > 30*time.Second {
			requestCount = 0
			windowStart = time.Now()
		}

		requestCount++

		if requestCount > 5 {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}

		fmt.Println("Request:", requestCount)

		next.ServeHTTP(w, r)
	})
}

func profileMiddleware(next http.Handler) http.Handler{
	
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){

		fmt.Println("middleware profile start")

		header:=r.Header.Get("cmrohit")

		if header != "yadav" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return // Bahut important
		}

		
		next.ServeHTTP(w,r)
	

		fmt.Println("middleware profile end")
	})

}

func logging(next http.Handler) http.Handler{
	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		fmt.Println("========== REQUEST START ==========")
		fmt.Println("Method :", r.Method)
		fmt.Println("Host   :", r.Host)
		fmt.Println("Path   :", r.URL.Path)
		fmt.Println("Start  :", start.Format("15:04:05"))

		next.ServeHTTP(w, r)

		end := time.Now()

		fmt.Println("End    :", end.Format("15:04:05"))
		fmt.Println("Duration :", end.Sub(start))
		fmt.Println("=========== REQUEST END ===========")
		fmt.Println()
})
}

func profile(w http.ResponseWriter,r *http.Request){
	
	w.Write([]byte("Welcome to Profile"))

}
func login(w http.ResponseWriter,r *http.Request){
	
	w.Write([]byte("Welcome to Login page"))

}
func home(w http.ResponseWriter,r *http.Request){
	
		w.Write([]byte("Welcome to Home page"))

}
func contact(w http.ResponseWriter,r *http.Request){
	
		w.Write([]byte("Welcome to contactpage"))

}

func main() {

	mux:=http.NewServeMux()

	mux.HandleFunc("/contact",contact)
	mux.Handle("/home",http.HandlerFunc(home))
	mux.Handle("/login",http.HandlerFunc(login))
	mux.Handle("/profile",profileMiddleware(http.HandlerFunc(profile)))
	

	mainmultiplexing:=rateLimiter(logging(mux))



	http.ListenAndServe(":3000",mainmultiplexing)
}