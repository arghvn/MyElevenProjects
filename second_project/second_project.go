package main

// go_movies_crud
// gorillamux We install this package and use it

import(
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

// See Redemy
// The reason that libraries give an error immediately after definition is that we imported them but they were not used

// define two structs , types = movie and director


// movies and directors are related, meaning that every film has a director
// movies can have other features that we define in its structure. 

type Movie struct{
     ID string "json:ID"
	 Isbn string "json:Isbn"
	 title string "json:title"
	 Director *Director "json:Director"


	 // Isbn is the unique index of each movie
	 // title is the The title of each movie

	// *Director is a pointer that means if i create a struct call director, it will be associated to movie struct 

// The personal information of each director is placed in the relevant structure
// Note that to avoid complexity, we define that each film has only one director.

}




type Director struct{
	Firstname string "json:Firstname"
	Lastname string "json:Lastname"


}


 // We need to define a variable that is a movie slice.

 var movies []Movie  

 func main(


	
	r := mux.NewRouter()
	r.HandleFunc()
	// r.HandleFunc() we will have five functions here

	// r is a router

	// NewRouter is a function inside mux library
 )