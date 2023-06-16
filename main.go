// Test package for improve my damn skills
package main

// Comment for imported packages
import (
	"fmt"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
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
	m := make(marks, 3)
	m["alex"] = 4

	// call print comment
	fmt.Println(m)

	_, err := m.byName("sdf")
	log.Error(err)
	// one more comment for another print
	fmt.Println("Starting...", 34)
	http.HandleFunc("/", handler)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("Could not start a server")
	}

}

// marks comment for type marks
type marks map[string]int

// byName comment for method from marks struct
func (m marks) byName(name string) (int, error) {
	mark, ok := m[name]
	if !ok {
		err := errors.New("undefined student")
		return 0, errors.Wrap(err, "byName method")
	}

	return mark, nil
}

func testPrint() string {
	return "TEST..."
}
