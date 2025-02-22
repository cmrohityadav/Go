package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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
var courses = []Course

// middleware or helper

func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
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

func getOnecourse(res http.ResponseWriter, req *http.Request) {
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
