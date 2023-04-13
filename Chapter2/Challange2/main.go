package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type book struct {
    ID     int
    Title  string
    Author string
    Desc   string
}

var books = []book{
	{ID: 1, Title: "Golang", Author: "Gopher", Desc: "A book for Go"},
	{ID: 2, Title: "Clean Code", Author: "Robert C. Martin", Desc: "A Handbook of Agile Software Craftsmanship"},
	{ID: 3, Title: "The Pragmatic Programmer", Author: "Andrew Hunt and David Thomas", Desc: "From Journeyman to Master"},
	{ID: 4, Title: "Design Patterns", Author: "Erich Gamma, Richard Helm, Ralph Johnson, John Vlissides", Desc: "Elements of Reusable Object-Oriented Software"},
	{ID: 5, Title: "Code Complete", Author: "Steve McConnell", Desc: "A Practical Handbook of Software Construction"},
}

var PORT = ":8080"

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/books", getBooks).Methods("GET")
    r.HandleFunc("/books/{id}", getBook).Methods("GET")
    r.HandleFunc("/books/create", createBook).Methods("POST")
    r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
    r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

    fmt.Println("Application is listening on port", PORT)
    http.ListenAndServe(PORT, r)
}

func getBooks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)

    for _, book := range books {
        if strconv.Itoa(book.ID) == params["id"] {
            json.NewEncoder(w).Encode(book)
            return
        }
    }

    http.Error(w, "Book not found", http.StatusNotFound)
}

func createBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var newBook book

    // membaca data JSON dari request body
    err := json.NewDecoder(r.Body).Decode(&newBook)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    newBook.ID = len(books) + 1
    books = append(books, newBook)

    json.NewEncoder(w).Encode(newBook)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)

    for index, b := range books {
        if strconv.Itoa(b.ID) == params["id"] {
            var updatedBook book

            // membaca data JSON dari request body
            err := json.NewDecoder(r.Body).Decode(&updatedBook)
            if err != nil {
                http.Error(w, err.Error(), http.StatusBadRequest)
                return
            }

            updatedBook.ID = b.ID
            books[index] = updatedBook

            json.NewEncoder(w).Encode(updatedBook)
            return
        }
    }

    http.Error(w, "Book not found", http.StatusNotFound)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)

    for index, book := range books {
        if strconv.Itoa(book.ID) == params["id"] {
            books = append(books[:index], books[index+1:]...)
            json.NewEncoder(w).Encode(book)
            return
        }
    }

    http.Error(w, "Book not found", http.StatusNotFound)
}
