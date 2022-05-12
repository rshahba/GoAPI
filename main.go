package main

import (
	"encoding/xml"
	"fmt"

	//"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Type    string `xml:"type"`
	Name    string `xml:"name"`
	ID      string `xml:"id"`
	Web     string `xml:"web"`
	Company string `xml:"company"`
}

type Users []User

func allUsers(w http.ResponseWriter, r *http.Request) {
	userId := mux.Vars(r)["id"]
	userName := mux.Vars(r)["name"]
	users := Users{
		User{Type: "admin", Name: "sara", ID: "22", Web: "Access Permitted", Company: "CIBC"},
		User{Type: "user", Name: "jason", ID: "55", Web: "Access Denied", Company: "XYZ"},
	}
	//var foundUser Users

	for i := 0; i < len(users); i++ {
		if (userId == users[i].ID) && (userName == users[i].Name) {
			fmt.Println("User ID match!")
			file, _ := xml.Marshal(users[i])
			w.Write([]byte(file))
			break
		}
		fmt.Println("User not found!")
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/users/{id}/{name}", allUsers).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	handleRequests()
}
