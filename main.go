package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	//hello world
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "welcome to the long dark teatime of the soul\n")
	}
	http.HandleFunc("/hello", helloHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}

}
