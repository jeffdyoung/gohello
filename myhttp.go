package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
func supGurl(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "SupGurl!")
	fmt.Println("Endpoint Hit: supGurl")
}
func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/sup", supGurl)
	log.Fatal(http.ListenAndServe("0.0.0.0:3000", nil))
}
