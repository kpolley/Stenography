package main

import (
	"fmt"
  "main/http"
)

func main() {
	fmt.Println("Initializing Endpoints")
  http.InitEndpoints()

  fmt.Println("Starting Server")
  http.StartServer()
}
