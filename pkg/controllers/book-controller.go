package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pfirulo2022/go-bookstore.git/pkg/models"
	"github.com/pfirulo2022/go-bookstore.git/pkg/utils"
)

var NewBook models.Book

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	booksDetails, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		booksDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		booksDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		booksDetails.Publication = updateBook.Publication
	}

	db.Save(&booksDetails)
	res, _ := json.Marshal(booksDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error al convertir el id")
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks, err := models.GetAllBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(newBooks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(res)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error writing response: %v", err), http.StatusInternalServerError)
	}
}

// func GetBook(w http.ResponseWriter, r *http.Request) {
// 	newBooks := models.GetAllBooks()
// 	res, _ := json.Marshal(newBooks)
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	_, err := w.Write(res)
// 	if err != nil {
// 		panic(err)
// 	}
// }

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error parsing id ", err)
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
