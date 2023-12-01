package main

import (
	"fmt"
	"net/http"
)

// func getMovies(w http.ResponseWriter, r *http.Request) {

// }

func main() {
	// http.HandleFunc("/movies", getMovies())
	// http.HandleFunc("/movies/", getMovie())
	// http.HandleFunc("/movies", createMovie())
	// http.HandleFunc("/movies/", updateMovie())
	// http.HandleFunc("/movies/", deleteMovie())

	var address = "localhost:8080"
	fmt.Printf("Server started at %s\n", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
