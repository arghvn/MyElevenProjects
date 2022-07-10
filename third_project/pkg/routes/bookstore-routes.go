package routes

import (
	"github.com/arghvn/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var registerBookStoreroutes = func(router *mux.Router) {
	router.Handlefunc("/book/", controllers.CreataBook).Methods("POST")
	router.Handlefunc("/book/", controllers.GETBook).Methods("GET")
	router.Handlefunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.Handlefunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.Handlefunc("/book/{bookId}", controllers.DeletBook).Methods("DELETE")
}
