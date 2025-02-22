package main

import (
	"math/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// model fro course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}
type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses = []Course{}

// middleware or helper

func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000", r))
}

//controllers - file

// serve home route
func serveHome(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("<h1>Welcome to API </h1>"))
}

func getAllCourses(res http.ResponseWriter, req *http.Response) {
	fmt.Println("Get all courses")
	res.Header().Set("content-Type", "application/json")
	json.NewEncoder(res).Encode(courses)

}

func getOneCourse(res http.ResponseWriter, req *http.Request) {
	fmt.Println("get one course")
	res.Header().Set("content-Type", "application/json")

	//grab id from request
	params := mux.Vars(req)

	//loop through courses, find matching id and return the response
	for index, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(res).Encode(course)
			return
		}
	}
	json.NewEncoder(res).Encode("No Course found with given id")
	return

}

func createOneCourse(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Create one course")
	res.Header().Set("content-Type", "application/json")

	//what if : body is empty
	if req.Body == nil {
		json.NewEncoder(res).Encode("Please send some data")

	}
	//what about if:{}
	var course Course
	_ = json.NewDecoder(req.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(res).Encode("No data inside JSON")

		return
	}

	//generate unique id, string
	rand.Seed(time.Now().UnixNano())
	course.CourseId=strconv.Itoa(rand.Intn(100))
	courses=append(courses,course)
	json.NewEncoder(res).Encode(course)
	return

}

func updateOneCourse(res http.ResponseWriter, req *http.Request){
	fmt.Println("update on course")
	res.Header().Set("content-Type", "application/json")

	//first: grab id from req
	params:=mux.Vars(req)

	//loop ,id,remove,add with my id
    for index,course:=range courses{
		if course.CourseId==params["id"]{
			courses=append(courses[:index],courses[index+1:]... )
			var course Course
			_=json.NewDecoder(req.Body).Decode(&course)
			course.CourseId=params["id"]
			courses=append(courses, course)
			json.NewEncoder(res).Encode(course)
			return


		}

	}


}

func deleteOneCourse(res http.ResponseWriter,req *http.Request){
	
}
