package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "post request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "name = %s\n", name)
	fmt.Fprintf(w, "address = %s\n", address)
}

func helloHandler(w http.ResponseWrite, r *http.Request) {
	// r = request (to server)... r is a pointer (pointing to request)
	// request is something that the user sends to the server and response is what the server sends back to the user.
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
		// alone return means we will return from the function
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
	// The fmt.Fprintf() function in Go language formats according to a format specifier and writes to w.
}

func main() {
	// We define our file server inside the main function.
	fileServer = http.FileServer(http.Dir("./static"))
	// Dir = directory
	// The handle function is required for the file server.
	http.Handle("/", http.FileServer)
	// FileServer is my route
	// A route generally refers to the route that a packet travels across a network.
	http.HandleFunc("/form", formHandler)
	// /form related to form.htmlfile
	http.HandleFunc("/Hello", helloHandlerfunction)
	// /Hello related to print just Hello
	// both of them must have handler function
	// at which time our server has allow to connect ? we will print this time
	fmt.Printf("starting server at port 8080?\n")
	if erro := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
		// nil means the error can be exist
		// != nil  is error handling and says if error not exist do next task
		// What does log fatal do in Go?
		// Every log message is output on a separate line
	}
	// The Print() function  prints a list of variables  with default formatting.
	// The Printf() function allows you to specify the formatting using a format template.
}

// first run
