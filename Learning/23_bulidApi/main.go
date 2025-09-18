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
	CoursePrice float64     `json:"price"`
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
	//seeding fake db
	courses = append(courses, 
		Course{CourseId: "1", CourseName: "React JS", CoursePrice: 500, Author: &Author{FullName: "Hitesh Sir", Website: "hitesh@lco.com"}},
		Course{CourseId: "2", CourseName: "GoLang Basics", CoursePrice: 700, Author: &Author{FullName: "John Doe", Website: "john@go.dev"}},
		Course{CourseId: "3", CourseName: "Python for Beginners", CoursePrice: 400, Author: &Author{FullName: "Jane Smith", Website: "jane@python.org"}},
		Course{CourseId: "4", CourseName: "Node.js Masterclass", CoursePrice: 650, Author: &Author{FullName: "Mark Lee", Website: "mark@nodejs.com"}},
		Course{CourseId: "5", CourseName: "Django Fullstack", CoursePrice: 550, Author: &Author{FullName: "Alice Brown", Website: "alice@django.dev"}},
	)
   
	r.HandleFunc("/",serveHome).Methods("GET")
	r.HandleFunc("/courses",getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}",getOneCourse).Methods("GET")
	r.HandleFunc("/course",createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}",updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}",deleteOneCourse).Methods("DELETE")

	//listen to a port
	fmt.Println("Server running on port 4000")
	log.Fatal(http.ListenAndServe(":4000", r))
}

//controllers - file

// serve home route
func serveHome(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("<h1>Welcome to API </h1>"))
}

func getAllCourses(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Get all courses")
	res.Header().Set("content-Type", "application/json")
	json.NewEncoder(res).Encode(courses)

}

func getOneCourse(res http.ResponseWriter, req *http.Request) {
	fmt.Println("get one course")
	res.Header().Set("Content-Type", "application/json")

	//grab id from request
	params := mux.Vars(req)

	//loop through courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(res).Encode(course)
			return
		}
	}
	res.WriteHeader(http.StatusNotFound)
	json.NewEncoder(res).Encode(map[string]string{"error": "No course found with given ID"})
	
	

}

func createOneCourse(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Create one course")
	res.Header().Set("Content-Type", "application/json")

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
	fmt.Println("created course : ",course)
	courses=append(courses,course)
	json.NewEncoder(res).Encode(course)
	

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
	fmt.Println("update on course")
	res.Header().Set("content-Type", "application/json")

	params:=mux.Vars(req)

	var course Course

	course.CourseId=params["id"]

	for index,c:=range courses{
		if c.CourseId==course.CourseId{
			courses=append(courses[:index],courses[index+1:]... )
			break
		}
	}

}
