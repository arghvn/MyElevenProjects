package main

// go_movies_crud
// gorillamux We install this package and use it

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

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
 json.NewEncoder(w).Encode(movies)
}
// r is the request that the user sends
// mux is the package that we are using for creating this routes and vars is inside mux which help us to get access to the request(params)
// All the videos are on one slide and we use them in the loop to check them one by one
// if we get a particular movie from the slice of the movies , we must find it and encode into json and send it to the front end or to postman

func getmovie(w http.ResponseWriter, r *http.request){
w.Header().set("Content-type", "application/json")
params := mux.vars(r)
// using for loop without index
for _, item := range movies {
	if item.ID == params["id"]
	    json.NewEncoder(w).Encode(item)
}
}
// inside params i have passed an id that use for compare each element in slice(movies) and if item.ID was equal to id ,do next command above(json...)
func createmovie(w http.ResponseWriter, r *http.request){
	w.Header().Set("content-type", "application/json")
	var movie Movie
	// a variable by name movie and type Movie
	_ = json.NewDecoder(r.body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	sfon.NewEncoder(w).Encode(movie)
	// we use rand for assign id to movie created
	// Intn is a function in random pacakge for between 1 and a number in ()
    // Itoa is a funcion
	// using strconv to convert the id created in rand to string
	func UpdateMovie (w http.http.ResponseWriter, r *http.Request){
		// set json content type
		w.Header().set("Content-Type", "application/json")
		// params
		params := mux.Vars(r)
		// loop over the movies ,range 
		// delete the movie with the id that you have been sent 
	// add a new movie - the movie that we send in the body of postman
	 for index, item := range movies{
		if item.ID == params["id"]{
		movies = append(movies[:index], movies[indeindex+1:]...)
		var movie Movie
		_*json.NewDecoder(r.Body).Decode(&movie)
		movie.ID = params["id"]
		movies = append(movies, movie)
		json.NewEncoder(w).Encode(movie)
        return
	 }
	}
	}
	// in UpdateMovie function we will pass an id, we will get id in params in mux and vars
	// in UpdateMovie function we at first delete a movie and then create a new id for it
	// the stages in this function instead of really updating(instead of updating a element in databasewe are just 
	// deleting it and then just adding a new element from the body that you are sending) : 
	// set json content type , params , loop over the movies ,range , delete the movie with the id that you have been sent ,
	// add a new movie - the movie that we send in the body of postman
	
}
// while creating a movie we will sen sth in the body ,we will send an entire movie to the body in postman and we must decode body
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

// By entering go build and then go run in the terminal, the server will now start up and ask us for access to the firewall
// postman intro