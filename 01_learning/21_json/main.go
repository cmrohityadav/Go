package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
type Course struct{
	Name string `json:"coursename"`
	Price int
	Platform string
	Password string `json:"-"` //just
	Tags [] string

}

type Student struct {
	Name  string `json:"name"`
	Class int    `json:"class"`
	DOY   int    `json:"doy"`
}

// Define a struct to match the full JSON response
type ApiResponse struct {
	Message string    `json:"message"`
	Data    []Student `json:"data"`
}

func main() {
	fmt.Println("Welcome to JSON")
	// EncodJson()
	// DecodeJson()
	handleApi()
}

func  EncodJson(){
	myCourse:=[]Course{
		{"ReactJs",500,"youtube.com","123456",[]string{"reactjs","javascript","web development"}},
		{"C++",1000,"youtube.com","123456",[]string{"system","DSA"}},
		{"JavaScript",1500,"youtube.com","123456",nil},
	}

	finalJson,err:=json.MarshalIndent(myCourse,"","\t")
	if err!=nil{
		panic(err)

	}
	fmt.Printf("%s\n",finalJson)

}

func DecodeJson(){
	jsonDataFromWeb:=[]byte(`
	{
                "coursename": "ReactJs",  
                "Price": 500,
                "Platform": "youtube.com",
                "Tags": [
                        "reactjs",        
                        "javascript",     
                        "web development" 
                ]
        }`)

		var myCourse Course
		checkValid:=json.Valid(jsonDataFromWeb)
		if checkValid {
			fmt.Println("JSON was Valid")
			json.Unmarshal(jsonDataFromWeb,&myCourse)

			fmt.Printf("%#v\n",myCourse)

		}else{
			fmt.Println("Json was not valid")
		}

		//some cases where you just want to add data to key value
		var myOnlineData map[string]interface{}
		json.Unmarshal(jsonDataFromWeb,&myOnlineData)	
		fmt.Printf("%#v",myOnlineData)
		for ky,vl:=range myOnlineData{
			fmt.Printf("\nKey is %v and value is %v and Type is %T",ky,vl,vl)
		}

}
func handleApi(){
	 myuri:="http://localhost:8000/apijson"
	responseFromApi,err:=http.Get( myuri)
	if err!=nil{
		panic(err)
	}

	defer responseFromApi.Body.Close()
	fmt.Println("Printing response only : ",responseFromApi)
	fmt.Println("Printing response Body : ",responseFromApi.Body)
	responseInByte,_:=io.ReadAll(responseFromApi.Body)

	fmt.Println("In string format : ",string(responseInByte))

	fmt.Printf("\nHandling api response\n\n")


	// Create a variable to store the decoded response
	var apiResponseMsgData ApiResponse

	// Unmarshal (decode) JSON into the struct
	err = json.Unmarshal(responseInByte, &apiResponseMsgData)
	if err != nil {
		panic(err)
	}

	// Print the structured data
	fmt.Println("Printing api response after unmarshal")
	fmt.Println(apiResponseMsgData)
	fmt.Println("Message:", apiResponseMsgData.Message)
	for i, student := range apiResponseMsgData.Data {
		fmt.Printf("Name: %s | Class: %d | Year of Birth: %d\n",
			student.Name, student.Class, student.DOY)
		fmt.Println("i : ",i)
	}



}