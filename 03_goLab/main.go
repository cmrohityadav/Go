package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
type User struct{
	Username string `json:"username"`
	Password string  `json:"password"`
	Opt      int  `json:"opt"`
}
func getUser(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Name", "Yadav")

	user:=User{
		Username: "cmrohityadav",
		Password: "rohit",
		Opt: 123,
	}

	jsonData,err:=json.Marshal(user)
	if err!=nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)

}

func createUser(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var data map[string]any

	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(data)

	fmt.Fprintf(w, "Received Successfully")

}
func main() {

	http.HandleFunc("/getuser",getUser)

	err:=http.ListenAndServe(":3000",nil)
	if err!=nil{
		fmt.Printf("Error while booting server %s",err)
	}
}