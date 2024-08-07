// ============== //
// SERVER SERVICE //
// ============== //

package main

import (
	"fmt"
	"io"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("helloHandler called")
	io.WriteString(w, "Hello, World!\n")
}

func rickHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("rickHandler called")
	io.WriteString(w, "Hello, World!\n")
}

func sayHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("sayHandler called")
	io.WriteString(w, "Hello, World!\n")
}

func searchBooksHandler(w http.ResponseWriter, r *http.Request) {
	response := BookResponse{
		Status: "success",
		Books: []Book{
			{Title: "Go Programming", Author: "John Doe"},
		},
	}

	fmt.Fprintln(w, response)
}

type Book struct {
	Title  string
	Author string
}

type BookResponse struct {
	Status string
	Books  []Book
}

// func rootHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("rootHandler called")
// 	io.WriteString(w, "You are on main page!\n")
// }

func main() {
	simplest()
	// withMux()
	// coolerServer()

}

func simplest() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/rick", rickHandler)
	http.HandleFunc("/say", sayHandler)
	http.HandleFunc("/searchBooks", searchBooksHandler)

	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}

// GET /hello
// 	GET /api/v1/rick
// 	GET /api/v1/say (+query "name")
// 	GET /api/v1/searchBooks (+query "title" and "author" )
// 	answer:
// 		type Book struct {
//     		title string`
//     		author string
// 		}
// 		type BookResponse struct {
//    			status string
//     		books  []Book
// 		}
// 	POST /api/v1/register
// name (string)
// password (string)
// 	POST /api/v1/calculate
// first_num (float64)
// second_num (float64)
// action (string)
// addition
// subtraction
// multiplication
// division
// 	POST /api/v1/translateText
// text (string)
// sourceLanguage (string) f.e. "en"
// targetLanguage (string) f.e. "ua"
//     	answer:
// status (string)
// translatedText (string)
// func withMux() {
// 	mux := http.NewServeMux()
// 	// mux.HandleFunc("/", rootHandler)
// 	mux.HandleFunc("/hello", helloHandler)

// 	fmt.Println("Server started on port 8080")
// 	http.ListenAndServe(":8080", mux)
