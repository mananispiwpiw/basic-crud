package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
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

func handlerMovie(w http.ResponseWriter, r *http.Request) {

	// Check if request Method type is GET
	if r.Method == http.MethodGet {
		// And if it is the correct method, it will continue to here
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(movieDatabase)

	} else if r.Method == http.MethodPost { // Check if request method type is POST
		// Decode the JSON request body into a Movie struct
		var newMovie Movie
		err := json.NewDecoder(r.Body).Decode(&newMovie)
		if err != nil {
			http.Error(w, "ERROR decoding JSON", http.StatusBadRequest)
			return
		}
		// Generate ID
		newMovie.ID = strconv.Itoa(rand.Intn(200))
		// Add the new movie to the existing database
		movieDatabase[newMovie.ID] = newMovie
		// Return with the newest Database as JSON
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(movieDatabase)

	} else {
		http.Error(w, "Method not allowed!", http.StatusMethodNotAllowed)
		return
	}
}

func handlerMovies(w http.ResponseWriter, r *http.Request) {
	// Extract the ID from the URL
	re := regexp.MustCompile(`/movies/(\d+)`)
	match := re.FindStringSubmatch(r.URL.Path)
	if len(match) != 2 {
		http.NotFound(w, r)
		return
	}
	movieID := match[1]

	// Look up the movie in the database
	movie, found := movieDatabase[movieID]
	if !found {
		http.NotFound(w, r)
		return
	}

	// Check if request Method type is GET
	if r.Method == http.MethodGet {
		// Return the response as JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(movie)
	} else if r.Method == http.MethodDelete { // Check if request Method type is GET
		// Perform deletion
		delete(movieDatabase, movieID)

		// Return the response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(movieDatabase)
	} else if r.Method == http.MethodPut { // Check if the request method type is PUT
		// Decode the JSON request body into a Movie struct
		var newMovie Movie
		err := json.NewDecoder(r.Body).Decode(&newMovie)
		if err != nil {
			http.Error(w, "ERROR decoding JSON", http.StatusBadRequest)
			return
		}
		// Perform updating
		movieDatabase[movieID] = newMovie
		// Return the response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(movieDatabase)
	} else {
		http.Error(w, "Method not allowed!", http.StatusMethodNotAllowed)
		return
	}
}

func main() {
	mux := http.NewServeMux()

	// Routes
	mux.HandleFunc("/movies", handlerMovie)
	mux.HandleFunc("/movies/", handlerMovies)

	var address = "localhost:8080"
	fmt.Printf("Server started at %s\n", address)
	err := http.ListenAndServe(address, mux)
	if err != nil {
		fmt.Println(err.Error())
	}
}
