package main
import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", sayHello)

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world demo16")
}
