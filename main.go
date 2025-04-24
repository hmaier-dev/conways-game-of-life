package main

import (
	"fmt"
	"net/http"
	"log"
)

func main() {
  const port = "8080"
  
  host := fmt.Sprintf("0.0.0.0:%s", port)
  srv := NewServer()
	log.Printf("Starting tool on %s \n", host)
	err := http.ListenAndServe(host, srv.Router)
	if err != nil {
		log.Fatal("cannot listen and server", err)
	}
}
