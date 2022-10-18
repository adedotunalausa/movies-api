package movies_api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	ID       string    `json:"id,omitempty"`
	Isbn     string    `json:"isbn,omitempty"`
	Title    string    `json:"title,omitempty"`
	Director *Director `json:"director,omitempty"`
}

type Director struct {
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
}

var movies []Movie

func main() {
	router := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "4895834", Title: "Yes Lord", Director: &Director{Firstname: "Mercy", Lastname: "Grace"}})
	movies = append(movies, Movie{ID: "2", Isbn: "5787539", Title: "Sweet Jesus", Director: &Director{Firstname: "Harry", Lastname: "Fin"}})

	router.HandleFunc("/movies", getMovies).Methods(http.MethodGet)
	router.HandleFunc("/movie/{id}", getMovie).Methods(http.MethodGet)
	router.HandleFunc("/movies", createMovie).Methods(http.MethodPost)
	router.HandleFunc("/movies/{id}", updateMovie).Methods(http.MethodPut)
	router.HandleFunc("/movies/{id}", deleteMovie).Methods(http.MethodDelete)

	fmt.Println("Server started on port 8950")
	log.Fatal(http.ListenAndServe(":8950", router))
}

func getMovies(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(responseWriter).Encode(movies)
	if err != nil {
		log.Fatalf("Error while fetching movies: %v \n", err)
		return
	}
}

func getMovie(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for _, item := range movies {
		if item.ID == params["id"] {
			err := json.NewEncoder(responseWriter).Encode(item)
			if err != nil {
				log.Fatalf("Error while getting movie with ID %v: %v \n", item.ID, err)
				return
			}
			return
		}
	}
}

func createMovie(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(request.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	err := json.NewEncoder(responseWriter).Encode(movie)
	if err != nil {
		log.Fatalf("Error while deleting movie: %v \n", err)
		return
	}
}

func updateMovie(responseWriter http.ResponseWriter, request *http.Request) {

}

func deleteMovie(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	err := json.NewEncoder(responseWriter).Encode(movies)
	if err != nil {
		log.Fatalf("Error while deleting movie: %v \n", err)
		return
	}
}
