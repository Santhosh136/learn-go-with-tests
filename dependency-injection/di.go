package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello, %s", name)
}

// print to standard output
// func main() {
// 	Greet(os.Stdout, "Santhosh")
// }

func GreetingHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Web")
}

func main() {
	http.ListenAndServe(":3000", http.HandlerFunc(GreetingHandler))
}
