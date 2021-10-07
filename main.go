package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func handleRequest(rspWriter http.ResponseWriter, request *http.Request) {
	for k, v := range request.Header {
		fmt.Printf("%s: %s\n", k, strings.Join(v, ""))
		rspWriter.Header().Add(k, strings.Join(v, ""))
	}
	rspWriter.Header().Add("version", os.Getenv("VERSION"))
	fmt.Fprintf(rspWriter, "You have gotted!\r")
}

func main() {
	http.HandleFunc("/", handleRequest)
	err := http.ListenAndServe(":9998", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
