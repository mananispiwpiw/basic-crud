package main

import (
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

// func getMovies(w http.ResponseWriter, r *http.Request) {

// }

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
	//Struct Test
	fmt.Println("Movie Name: ", movie.Name)
	fmt.Println("Movie Director: ", movie.Director.FirstName+" "+movie.Director.LastName)

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
