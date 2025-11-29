package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod in go lang")
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome)
	log.Fatal(http.ListenAndServe(":4000", r))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello from Go lang</h1>"))
}
