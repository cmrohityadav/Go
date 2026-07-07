package main

import (
	"fmt"
	"net/http"
)

func getInfo(w http.ResponseWriter,r *http.Request){
	query:=r.URL.Query() // return : type Values map[string][]string
	fmt.Println(query)
    // map[lang:[Go Cplusplus] middle:[Rohit] mobile:[1234567890] name:[Rahul] surname:[Yadav]]
	// query["key"] → saari values ([]string)
    // query.Get("key") → sirf pehli value (string)

	
	// -------------------- Get() --------------------
	name := query.Get("name") //return:  string
	middle := query.Get("middle")
	surname := query.Get("surname")
	mobile := query.Get("mobile")

	fmt.Fprintf(w,"my Name %s",name)
	fmt.Fprintf(w,"my middle name is %s and surname is %s",middle,surname)
	fmt.Fprintf(w,"my Mobile number %s",mobile)

	result:=fmt.Sprintf("\nname: %s\n Middle: %s\n Surname: %s\n Mobile: %s\n",name,middle,surname,mobile)


	w.Write([]byte(result))

	w.Write([]byte("hello"))

	// -------------------- Set() --------------------
	query.Set("mobile","9988774455") // Sets a value. If the key already exists, it replaces all existing values
	fmt.Fprintf(w, "After Set() Mobile : %s\n", query.Get("mobile"))

	// -------------------- Add() --------------------
	// Add() : It does not replace existing values
	query.Add("lang", "Python")
	query.Add("lang", "JavaScript")
	fmt.Fprintf(w, "After Add() Languages : %v\n", query["lang"])

	// -------------------- Del() --------------------
	// Del() : Deletes a key completely
	query.Del("surname")
	fmt.Fprintf(w, "After Del() surname : %q\n", query.Get("surname"))


	// -------------------- Has() --------------------
	// Has() : Checks whether a key exists
	fmt.Fprintf(w, "Has name? %v\n", query.Has("name"))
	fmt.Fprintf(w, "Has surname? %v\n", query.Has("surname"))
	
	// -------------------- Encode() --------------------
	// Encode() : Converts the map back into a URL query string
	encoded := query.Encode()

	fmt.Fprintf(w, "\nEncoded Query:\n%s\n", encoded)

	// -------------------- Access All Values --------------------
	fmt.Fprintf(w, "\nAll Languages : %v\n", query["lang"])
    

	// -------------------- Check Key Exists --------------------
	if langs, ok := query["lang"]; ok {
		fmt.Fprintf(w, "Lang exists : %v\n", langs)
	}

	// -------------------- Iterate --------------------
	fmt.Fprintf(w, "\nIterating Query Params:\n")

	for key, values := range query {
		fmt.Fprintf(w, "%s = %v\n", key, values)
	}
}
func main() {

	http.HandleFunc("/getinfo",getInfo)

	err:=http.ListenAndServe(":3000",nil)
	if err!=nil{
		fmt.Println("Error while booting up server : ",err)
	}
}