package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	listenPortKey = "PORT"
)

func handler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "hello world!")
}

func main() {
	http.HandleFunc("/", handler)
	listenPort := os.Getenv(listenPortKey)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", listenPort), nil); err != nil {
		log.Fatalf("failed to listen and serve: %v", err)
	}
	log.Print("listen...")
}
