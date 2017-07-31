package main

import
(
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "encoding/json"
)

type Person struct {
	ID 			string 		'json:"id,omitempty'
	FirstName 	string 		'json:"firstname,omitempty"'
	LastName 	string 		'json:"lastname,omitempty"'
	Address 	*Address	'json:"address, omitempty'
}

type Address struct {
	City 	string 'json:"city,omitempty'		
	State 	string 'json:"state,omitempty'		
}

var people []Person

func GetPersonEndpoint(w http.ResposeWriter, req *http.Request) {

}

func GetPeopleEndpoint(w http.ResposeWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func UpdatePersonEndpoint(w http.ResposeWriter, req *http.Request) {

}

func DeletePersonEndpoint(w http.ResposeWriter, req *http.Request) {

}

func main() {
//Istanza del router dell'applicazione
	router := mux.NewRouter()
	people := append(people, Person{ID: "1", Firstname: "Michele", Lastname:"Maresca", Address:&Address{City:"Messina", State:"Italy"}})
	people := append(people, Person{ID: "2", Firstname: "Alessio", Lastname:"Maresca", Address:&Address{City:"Messina", State:"Italy"}})
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePeopleEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePeopleEndpoint).Methods("DELETE")
//http oggetto che gestice il server
	log.Fatal(http.ListenAndServe(":12345", router))
}