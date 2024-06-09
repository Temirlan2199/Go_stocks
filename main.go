package main

import (
	"fmt"
	"go-postgres-yt/router"
	"log"
	"net/http"

	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

func main() {
	r := router.Router()
	fmt.Println("Server is running")

	log.Fatal(http.ListenAndServe(":8080", r))
}
