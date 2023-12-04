package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
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

// Make map, for database mocking
var movieDatabase = map[string]Movie{
	"123": {
		ID:    "123",
		Name:  "The conjurinng",
		Genre: "Horror",
		ISAN:  "0000-0000-E5F-0000-2-0000-0000-K",
		Director: &Director{
			ID:        "12",
			FirstName: "Christopher",
			LastName:  "Nolan",
		},
	},
	"124": {
		ID:    "124",
		Name:  "The Insidious",
		Genre: "Horror",
		ISAN:  "0000-0004-2E5A-0000-8-000-0100-A",
		Director: &Director{
			ID:        "4",
			FirstName: "Adam",
			LastName:  "Warlock",
		},
	},
}

//var movies []movieDatabase

func getMovies(w http.ResponseWriter, r *http.Request) {
	//Check if request Method type is not GET
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed!", http.StatusMethodNotAllowed)
		return
	}
	//And if it is the correct method, it will continue to here
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movieDatabase)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	//Check if request Method type is not GET
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed!", http.StatusMethodNotAllowed)
		return
	}
	//Extract the ID from the URL
	re := regexp.MustCompile(`/movies/(\d+)`)
	match := re.FindStringSubmatch(r.URL.Path)
	if len(match) != 2 {
		http.NotFound(w, r)
		return
	}
	movieID := match[1]

	//Look up the movie in the database
	movie, found := movieDatabase[movieID]
	if !found {
		http.NotFound(w, r)
		return
	}

	//Return the response as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}

func main() {
	http.HandleFunc("/movies", getMovies)
	http.HandleFunc("/movies/", getMovie)
	// http.HandleFunc("/movies", createMovie)
	// http.HandleFunc("/movies/", updateMovie)
	// http.HandleFunc("/movies/", deleteMovie)

	var address = "localhost:8080"
	fmt.Printf("Server started at %s\n", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
