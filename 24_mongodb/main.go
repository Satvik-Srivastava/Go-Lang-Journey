package main

import (
	"fmt"
	"log"
	"mongo_db/router"
	"net/http"
)

func main() {
	fmt.Println("MONGODB API")
	r := router.Router()
	fmt.Println("Sever is running....")

	log.Fatal(http.ListenAndServe(":5173", r))
	fmt.Println("Listening at PORT: 5173")
}
