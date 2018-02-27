package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Idea struct {
	ID          int    `bson:"_id" json:"id"`
	Description string `bson:"description" json:"description"`
}

func AllIdeas(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Future home of /ideas GET")
}

func FindIdea(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Future home of /ideas/:id GET")
}

func CreateIdea(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Future home of /ideas POST")

}

func UpdateIdea(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Future home of /ideas/:id PUT")
}

func DeleteIdea(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Future home of /ideas/:id DELETE")
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/ideas", AllIdeas).Methods("GET")
	r.HandleFunc("/ideas", CreateIdea).Methods("POST")
	r.HandleFunc("/ideas/{id}", UpdateIdea).Methods("PUT")
	r.HandleFunc("/ideas/{id}", DeleteIdea).Methods("DELETE")
	r.HandleFunc("/ideas/{id}", FindIdea).Methods("GET")

	http.ListenAndServe(":8080", r)
}
