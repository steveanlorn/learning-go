package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/steveanlorn/learning-go/testing/unit_test/05_test_internal_endpoint/handlers"
)

func main() {
	log.Println("Running server at 7070")
	if err := http.ListenAndServe(":7070", nil); err != nil {
		log.Printf("Could not run server at 7070: %v", err)
		os.Exit(1)
	}
}
