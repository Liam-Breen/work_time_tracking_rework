package main

import (
	"fmt"
	"log"
	"net/http"
)

func startMessage() string {
	return "STARTING SERVER..."
}

func main() {
	fmt.Println(startMessage())

	handler := http.HandlerFunc(UserServer)
	if err := http.ListenAndServe(":6969", handler); err != nil {
		log.Fatalf("could not listen on port 6969 %v", err)
	}
}
