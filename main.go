package main

import
(
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "encoding/json"
)

type Person struct {
	ID 			string 		`json:"id,omitempty"` 
	FirstName 	string 		`json:"firstname,omitempty"`
	LastName 	string 		`json:"lastname,omitempty"`
	Address 	*Address	`json:"address, omitempty`
}

type Address struct {
	City 	string `json:"city,omitempty"`		
	State 	string `json:"state,omitempty"`		
}

var people []Person

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range(people) {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
		}
	}
}

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {

}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {

}

func main() {
//Istanza del router dell'applicazione
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", FirstName: "YourName", LastName:"YourName", Address:&Address{City:"YourCity", State:"YourCity"}})
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
//http oggetto che gestice il server
	log.Fatal(http.ListenAndServe(":12345", router))
}