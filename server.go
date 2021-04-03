package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func GreetingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "bye")
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	// type conversion which have a method of ServerHTTP()
	if err := http.ListenAndServe(":"+port, http.HandlerFunc(GreetingHandler)); err != nil {
		log.Fatal(err)
	}
}
