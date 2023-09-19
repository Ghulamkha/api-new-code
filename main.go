package main

import (

"encoding/json"
"net/http"
"strconv"
"github.com/gorilla/mux"


)

var profile []profile= []profile{}

type User struct
{
Firstname string 'json: "firstname"'
Lastname string 'json: "lastname"'
Email string 'json: "email"'
}
type Profile struct {
	Department string 'json: "department"'
	Designation string 'json: "designation"'
	Employee user '"json: employee"'
}

func addItem(q http.ResponseWriter, r *http.Request)
{
var newProfile Profile
json.NewDecoder(r.body).Decode(&newProfile)
w.Header().Set("Content-Type", application/json)
profiles -append(profiles, newProfile)
json.Newencoder(q).Encode(profiles)
}

  func main() {
  router := mux.NewRouter()

  router.HandleFunc("/profiles", addItem).Methods("POST")

  http.listenAndServe(":5000", router)
  }