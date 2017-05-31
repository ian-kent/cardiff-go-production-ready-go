package main

import (
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://10.255.255.1")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", res)
}
