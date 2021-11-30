package main

import (
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)
//book struct
type Book struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title  string `json:"title"`
	Author *Author `json:"author"`
}

//Author struct
type  Author struct {
	Firstname string `json:"firstname`
	Lastname string `json:"lastname`
}

//init books variable as a slice books struct
var books []Book 
//get all books
func getBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
//single book
func getBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get params

	//find with id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}
//create new book
func createBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var book Book 
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000)) //mock id
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		for index, item := range books {
		if item.ID == params ["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book 
				_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params ["id"] //mock id
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params ["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
		json.NewEncoder(w).Encode(books)
}
func main(){
	//Init the mux router
	r := mux.NewRouter()

	//mock data
	books = append(books, Book{ID : "1", Isbn: "434546", Title: "Book One", Author: &Author{Firstname: "john", Lastname : "Doe"}})

	books = append(books, Book{ID : "2", Isbn: "434556", Title: "Book Two", Author: &Author{Firstname: "john", Lastname : "Smiith"}})

	//create route handler  / Endpoint
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
	
	
}