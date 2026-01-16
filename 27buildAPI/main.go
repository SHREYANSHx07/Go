package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CorseId     string  `json:"courseid"`
	CourseName  string  `json:"Name"`
	Author      *Author `json:"courseauthor"`
	CoursePrice int     `json:"Price"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"web"`
}

//DB

var courses []Course

func (c *Course) Isempty() bool {
	//return c.CorseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("welcome in API creation ")
	r := mux.NewRouter()
	courses = append(courses, Course{CorseId: "2",CourseName: "Golong",CoursePrice: 999, Author: 
	&Author{Fullname: "shreyansh Gupta",Website: "www.google.com"}})
	courses = append(courses, Course{CorseId: "3",CourseName: "JS",CoursePrice: 1999, Author: 
	&Author{Fullname: "shreyansh rao",Website: "www.youtube.com"}})
	courses = append(courses, Course{CorseId: "4",CourseName: "C++",CoursePrice: 599, Author: 
	&Author{Fullname: "sonali",Website: "www.instagram.com"}})
	courses = append(courses, Course{CorseId: "5",CourseName: "C",CoursePrice: 299, Author: 
	&Author{Fullname: "Vishal",Website: "www.facebook.com"}})

	r.HandleFunc("/",Servehome).Methods("GET")
	r.HandleFunc("/courses",Getall).Methods("GET")
	r.HandleFunc("/course/{id}",Getone).Methods("GET")
	r.HandleFunc("/course",create1course).Methods("POST")
	r.HandleFunc("/course/{id}",update1course).Methods("PUT")
	r.HandleFunc("/course/{id}",delete1course).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000",r))
}

func Servehome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to the cute person(IRaa)</h1>"))
}

func Getall(w http.ResponseWriter, r *http.Request) {
	fmt.Println("print all courses")
	w.Header().Set("cotent-type", "application/json")
	json.NewEncoder(w).Encode(courses)

}

func Getone(w http.ResponseWriter, r *http.Request) {
	fmt.Println("find the one course")
	w.Header().Set("cotent-type", "application/json")
	param := mux.Vars(r)

	for _, x := range courses {
		if x.CorseId == param["id"] {
			json.NewEncoder(w).Encode(x)
			return
		}
	}
	json.NewEncoder(w).Encode("No course with this valid ID")
	return

}

func create1course(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create the one create1course")
	w.Header().Set("cotent-type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.Isempty() {
		json.NewEncoder(w).Encode("No data in json")
		return
	}
	for _,i :=range courses{
		if i.CourseName == course.CourseName{
			json.NewEncoder(w).Encode("Course already exists")
			return
		}

	}

	rand.Seed(time.Now().UnixNano())
	course.CorseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func update1course(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update 1 course")
	w.Header().Set("cotent-type", "application/json")
	pramas := mux.Vars(r)
	// var x int
	// var course Course
	// _ = json.NewDecoder(r.Body).Decode(&course)
	// for _,i :=range courses{
	// 	if i.CorseId == course.CorseId{
	// 		x++
	// 	}

	// }
	// if(x==0){
	// 	fmt.Println("No data for this ID")
	// 	return
	// }

	
	for index, course := range courses {
		if course.CorseId == pramas["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CorseId = pramas["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No valid data for given ID")
	return 

}

func delete1course(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete 1 course")
	w.Header().Set("cotent-type", "application/json")
	pramas := mux.Vars(r)
	for index, course := range courses {
		if course.CorseId == pramas["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			fmt.Println("Deleted the course")
			break
		}
	}
	json.NewEncoder(w).Encode("no data with this ID")
	return

}
