package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// import (
// 	"log"
// 	"net/http"
// 	"github.com/gorilla/mux"
// 	"github.com/jizhu/gorm/dialects/mysql"
// 	"github.com/arghvn/go-bookstore/pkg/rotes"
// )

func main() {
	r := mux.NewRouter()
	routes.RegisterbookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
