package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// model of course or schema in fullstack
type Course struct {
	CourseId string `json:"courseId"`
	Name     string `json:"courseName"`
	Price    int    `json:"price"`
	// now we will inject our author field in this
	Author *Author // we are pointing not referencing
}

// model of the author
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// we will add some middleware which are basically helper

func (c *Course) isEmpty() bool {
	return c.Name == "" && c.CourseId == ""
}

func main() {
	fmt.Println("We are building an api")
	r := mux.NewRouter()
	r.HandleFunc("/", ServerHome)
	r.HandleFunc("/getcourse", getAllCourses)
	r.HandleFunc("/getonecourse", getOneCourses)
	r.HandleFunc("/updateonecourse", updateOneCourse)

}

var courses []Course // creating a fake db for now

// controller usually go in there own seperate files

// server Home page

func ServerHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<h1>This is the home page</h1>`))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all the courses")
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// get only one course
func getOneCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all the courses")
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	// loop over the courses and match the id with the id passed as request

	for _, val := range courses {
		if val.CourseId == params["id"] {
			json.NewEncoder(w).Encode(val)
			return
		}
	}
	json.NewEncoder(w).Encode("No matching course found")

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("updating one courses")
	w.Header().Set("Content-type", "application/json")
	// var coursera = []string{"react-js", "js", "swift", "python", "ruby","c++","java"}
	// fmt.Println(coursera)

	// first grab the id by params

	params := mux.Vars(r)

	// loop through the value, find the id, remove the value and add the another id

	for index, cors := range courses {
		if cors.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			json.NewDecoder(r.Body).Decode(&course)
		}
	}

}
