package main

import (
	"log"
	"net/http"
	"time"
)

var cli = http.Client{
	Timeout: time.Second,
	// Transport: &http.Transport{
	//     Dial: (&net.Dialer{
	//             Timeout:   30 * time.Second,
	//             KeepAlive: 30 * time.Second,
	//     }).Dial,
	//     TLSHandshakeTimeout:   10 * time.Second,
	//     ResponseHeaderTimeout: 10 * time.Second,
	//     ExpectContinueTimeout: 1 * time.Second,
	// }
}

func main() {
	res, err := cli.Get("http://10.255.255.1")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", res)
}
