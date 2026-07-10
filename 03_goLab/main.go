package main

import (
	
	"net/http"
)

func login(w http.ResponseWriter,r *http.Request){
	
	cookie:=&http.Cookie{
		Name: "session_id",
		Value: "secret_key",
		HttpOnly: true,
		Secure: true,
	}

	http.SetCookie(w,cookie)
	w.Write([]byte("Cookie set"))

}
func main() {

	
	http.HandleFunc("/login",login)

	http.ListenAndServe(":3000",nil)
}