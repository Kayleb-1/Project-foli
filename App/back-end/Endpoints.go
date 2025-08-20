package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const HOST_ADDR =  ip + ":" + port
const ip = "localhost"
const port = "8080"

const StatusOK = http.StatusOK



// Testing a basic server response
func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", GetTitle)
	log.Fatal(http.ListenAndServe(HOST_ADDR, r))

}


// GetTitle - Return hello world to client
func GetTitle(w http.ResponseWriter, r *http.Request){

	resp := "Hello"

	w.WriteHeader(StatusOK)
	w.Write([]byte(resp))
}