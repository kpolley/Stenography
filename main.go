package main

import (
	"fmt"
  "main/http"
)

func main() {
	fmt.Println("Hello, World!")
  http.InitEndpoints()
  http.StartServer()
}
