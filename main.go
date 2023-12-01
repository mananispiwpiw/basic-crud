package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Director struct {
	ID        string `json:"dirId"`
	FirstName string `json:"dirFirstName"`
	LastName  string `json:"dirLastName"`
}
type Movie struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Genre    string    `json:"genre"`
	ISAN     string    `json:"isan"`
	Director *Director `json:"director"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	//Check if request Method type is not GET
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed!", http.StatusMethodNotAllowed)
		return
	}
	//And if it is the correct method, it will continue to here
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func main() {
	//Director Instance
	director := Director{
		ID:        "12",
		FirstName: "Cristopher",
		LastName:  "Nolan",
	}
	//Movie Data
	movie := Movie{
		ID:       "123",
		Name:     "The Conjuring",
		Genre:    "Horror",
		ISAN:     "0000-0000-9E5F-0000-2-0000-0000-K",
		Director: &director,
	}
	movies = append(movies, movie)

	http.HandleFunc("/movies", getMovies)
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
