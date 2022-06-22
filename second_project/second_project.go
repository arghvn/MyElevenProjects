package main

// go_movies_crud
// gorillamux We install this package and use it

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// See Redemy
// The reason that libraries give an error immediately after definition
// is that we imported them but they were not used

// define two structs , types = movie and director

// movies and directors are related, meaning that every film has a director
// movies can have other features that we define in its structure.

type Movie struct {
	ID       string    "json:ID"
	Isbn     string    "json:Isbn"
	title    string    "json:title"
	Director *Director "json:Director"

	// Isbn is the unique index of each movie
	// title is the The title of each movie

	// *Director is a pointer that means if i create a struct call director,
	// it will be associated to movie struct

	// The personal information of each director is placed in the relevant structure
	// Note that to avoid complexity, we define that each film has only one director.

}

type Director struct {
	Firstname string "json:Firstname"
	Lastname  string "json:Lastname"
}

// We need to define a variable that is a movie slice.

var movies []Movie


func getMovies(w http.ResponseWriter, r *http.Request){
// we are passing the request a pointer
// we will send the request from our postman to this function and w is the responsewriter
// when we send a response from this function , it will be w
// we want to set the content type as json....it is struct
W.Header().set("Content-Type, application/json")
json.NewEncoder(w).Encode(movies)
// the response that we want to send to it , we want to encode it into json 
// and we want to pass the complete movies slice that we have in var movies []Movie
}
// when we want to delete a movie , we will pass an id of the movie
 func deletemovie(w http.ResponseWriter, r *http.Request){
	w.Header().set("content-type", "application/json")
	// params like id that pass by postman.it is part of the request and mux
	// by writting below line we have access to id 
	params := mux.vars(r)
	// range over movies...we got params ,params is the id that we sent as a request this funnction and get we access to particular movie 
	for index, item := range movies {

		if item.ID == params["id"]{
		movies = append(movies[:inedx], movies[index+1:]...)
		break
	}
 }
}
// r is the request that the user sends
// mux is the package that we are using for creating this routes and vars is inside mux which help us to get access to the request(params)
// All the videos are on one slide and we use them in the loop to check them one by one
// if we get a particular movie from the slice of the movies , we must find it and encode into json and send it to the front end or to postman

func getmovie(w http.ResponseWriter, r *http.request){
w.Header().set("Content-type", "application/json")
params := mux.vars(r)
}




func main() {
	r := mux.NewRouter()
	// at least two movie is enough to find how server working
	movies = append(movies, movie{ID: "1", Isbn:"438227", Title :"Movie one", Director : &Director(Firstname:"John", Lastname:"Doe")})
	movies = append(movies, movie{ID: "2", Isbn:"454555", Title :"Movie two", Director : &Director(Firstname:"Steve", Lastname:"smith")})
	// we use &Director because we want the reference object(the reference of the address of director) 
	// append movie to movies
	// r.HandleFunc() we will have five functions here
	// r is a router
	// NewRouter is a function inside mux library
	// first route is the movies route
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methos("POST")
	r.HandleFunc("/movies/{id}", updatemovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deletemovie).Methods("DELETE")
	// To set up the server, we say:
	fmt.Printf("starting server at port 8000\n")
	// i, going to log out system if the server doesnt start
	log.Fatal(http.ListenAndServe(":8000", r))
}

//  when we go to our postman and we will hit the server , when we hit the  API/movies and we want to get
//  all movies (in the beginning there wont be any movies so that we want to have at least two movies)
//  then in above we take movies struct, slice of movies
//  im going to append a couple of movies to it.
