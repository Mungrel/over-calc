package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Mungrel/over-calc/server"
)

func main() {
	fmt.Printf("serving on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", server.New()))
}
