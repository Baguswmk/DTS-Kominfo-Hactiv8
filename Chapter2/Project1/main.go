package main

import (
	"dts/Project1/database"
	"dts/Project1/models"

	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func main() {
	database.ConnectDB()

	router := mux.NewRouter()

	router.HandleFunc("/books", getAllBook).Methods("GET")
	router.HandleFunc("/books/{id}", getBookByID).Methods("GET")
	router.HandleFunc("/books", createBook).Methods("POST")
	router.HandleFunc("/books/{id}", updateBookByID).Methods("PUT")
	router.HandleFunc("/books/{id}", deleteBookByID).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))

}

func getAllBook(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()

	books := []models.Book{}
	db.Find(&books)
	db.Order("id desc")

	json.NewEncoder(w).Encode(books)
}

func getBookByID(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()

	book := models.Book{}

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		fmt.Println("Error converting id to int:", err)
		return
	}

	err = db.First(&book, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Book not Found")
			return
		}
		fmt.Println("Error finding book :", err)
	}

	json.NewEncoder(w).Encode(book)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()

	book := models.Book{}
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		fmt.Println("Error decoding request body:", err)
		return
	}

	err = db.Create(&book).Error

	if err != nil {
		fmt.Println("Error creating book", err)
		return
	}
	fmt.Println("Book created : ", book)

	json.NewEncoder(w).Encode(book)
}

func updateBookByID(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()
	book := models.Book{}

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		fmt.Println("Error converting id to int:", err)
		return
	}

	err = db.Model(&book).Where("id = ?", id).Updates(models.Book{Name_book: book.Name_book, Author: book.Author}).Error

	if err != nil {
		fmt.Println("Error updating book:", err)
	}
	response := map[string]string{"message": "Book Update successfully"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error creating JSON response:", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func deleteBookByID(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()

	book := models.Book{}

	id, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		fmt.Println("error deleting:", err.Error())
		return
	}

	err = db.Model(&book).Where("id = ?", id).Delete(&book).Error

	if err != nil {
		fmt.Println("Error deleting book:", err.Error())
	}	
	response := map[string]string{"message": "Book deleted successfully"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Error creating JSON response:", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}