// Test package for improve my damn skills
package main

// Comment for imported packages
import (
	"fmt"
	"net/http"
)

// handler describe hndlr func
func handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Helloworld")
	if err != nil {
		panic("Handle error")
	}
}

// main there is another comment for main func
func main() {

	// one more comment for another print
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
