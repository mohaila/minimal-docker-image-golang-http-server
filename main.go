package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Http!")
}

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}

	http.HandleFunc("/hello", helloHandler)

	log.Printf("Server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
