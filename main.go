package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"person"

	"github.com/gorilla/mux"
)

var people []person.Person

func main() {
	router := mux.NewRouter()

	// people := person.Person{}
	people = append(people, person.Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})
	fmt.Println(people)

	// people = append(people, Person{ID: "2", FirstÒname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})

	router.HandleFunc("/people", GetPeople).Methods("GET")
	// router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	// router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	// router.HandleFunc("/people/{id}", DeletePerson).Methods("DELEssssTE")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

// func GetPerson(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	for _, item := range people {
// 		if item.ID == params["id"] {
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode(&Person{})
// }

// func CreatePerson(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	var person Person
// 	_ = json.NewDecoder(r.Body).Decode(&person)
// 	person.ID = params["id"]
// 	people = append(people, person)
// 	json.NewEncoder(w).Encode(people)
// }

// func DeletePerson(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	for index, item := range people {
// 		if item.ID == params["id"] {
// 			people = append(people[:index], people[index+1:]...)
// 			break
// 		}
// 		json.NewEncoder(w).Encode(people)
// 	}
// }
