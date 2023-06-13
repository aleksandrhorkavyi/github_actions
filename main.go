package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Helloworld")
	if err != nil {
		panic("Handle error")
	}
}

func main() {

	fmt.Println("Starting...")
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("Could not start a server")
	}

}

func testPrint() string {
	return "TEST..."
}
