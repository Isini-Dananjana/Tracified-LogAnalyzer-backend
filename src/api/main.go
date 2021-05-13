package main

import (

	//"encoding/json"
	//"log"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"net/http"

	"math/rand"
	"strconv"

	"github.com/gorilla/mux"
)

//get All books

func getBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)

}

//get single book

func getBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	//getid
	params := mux.Vars(r) //get params

	//loop though book and find with id

	for _, item := range books {

		if item.ID == params["id"] {

			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})

}

//create a book

func craetebook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var book Book

	_ = json.NewDecoder(r.Body).Decode(&book)

	book.ID = strconv.Itoa(rand.Intn(100000)) // mock id -not safe

	books = append(books, book)
	json.NewEncoder(w).Encode(book)

}

//update

func updateBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //get params

	for index, item := range books {

		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)

			var book Book

			_ = json.NewDecoder(r.Body).Decode(&book)

			book.ID = params["id"] //

			books = append(books, book)
			json.NewEncoder(w).Encode(book)

			return

		}
	}
	json.NewEncoder(w).Encode(books)

}

//delete

func deleteBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //get params

	for index, item := range books {

		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)

			break
		}
	}
	json.NewEncoder(w).Encode(books)

}

//Book Struct(Model)

type Book struct {
	ID     string  `json:"id`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

//Author struct

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

//Init books var as a slice Book struct
var books []Book

type logFile struct {
	ID      string `json:"id`
	name    string `json:"name"`
	content string `json:"content"`
}

var logfile []logFile

func getLogFile(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(logfile)

}

func main() {

	//Init Router

	r := mux.NewRouter() //router var

	data, err := ioutil.ReadFile("file.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:")
	fmt.Println(string(data))

	var dataT = string(data)

	logfile = []logFile{

		logFile{ID: "1", name: "ldhu", content: dataT},
		
	}

	//router handlers//endpointa

	r.HandleFunc("/api/logs", getLogFile).Methods("Get")

	r.HandleFunc("/api/books", getBooks).Methods("Get")
	r.HandleFunc("/api/books/{id}", getBook).Methods("Get")
	r.HandleFunc("/api/books", craetebook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBooks).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBooks).Methods("DELETE")

	http.ListenAndServe(":8000", r)

}
//jdnkndk