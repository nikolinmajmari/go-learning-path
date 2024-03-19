package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(writer io.Writer, name string) error {
	_, err := fmt.Fprintf(writer, "Hello, %s", name)
	return err
}

func MyGreetWriter(w http.ResponseWriter, r *http.Request) {
	Greet(w, "World")
}

func main() {
	log.Fatal(http.ListenAndServe(":50001", http.HandlerFunc(MyGreetWriter)))
}
