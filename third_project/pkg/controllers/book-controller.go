package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/arghvn/go-bookstore/pkg/models"
	"github.com/arghvn/main.go/pkg/models"
	"github.com/gorilla/mux"
)

// import (
// 	"encoding/json"
// 	"fmt"
// 	"github.com/gorilla/mux"
// 	"net/http"
// 	"strconv"
// 	"github.com/arghvn/go-bookstore/pkg/utils"
// 	"github.com/arghvn/go-bookstore/pkg/models"
// )

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("content-type", "pkglocation/json")
	w.WriteHeader(httpStatusOk)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("content-type", "pkglocation/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// vars related to gorillamux
// in GO when using thing is not good we use _ , because if we define something and dont use it we will have an error

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// CreateBook is a data type

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	Vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("content-type", "pkglocation/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
