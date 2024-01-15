package main
import (

	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json:"id"`
	Language string `json:"language"`
	Name string `json:"name"`
	Director *Director `json:"director"`

}
type Director struct{
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}
var movies[]Movie

func getMovies(m http.ResponseWriter , r *http.Request){
m.Header().Set("Content-type", "application/json")
json.NewEncoder(m).Encode(movies)
}

func DeleteMovie(m http.ResponseWriter , r *http.Request){
	m.Header().Set("content-type" ,"application/json")
	params := mux.Vars(r)
	 for index, items :=range movies{
		if items.ID == params["id"]{
			movies=append(movies[:index], movies[index+1:]... )
			break
		}
	 }
 json.NewEncoder(m).Encode(movies)
}

func getMovieByID(m http.ResponseWriter , r *http.Request){
	m.Header().Set("content-type", "application/json")

	params:=mux.Vars(r)
	for _ , items:=range movies{
		if items.ID==params["id"]{
			json.NewEncoder(m).Encode(items)
			return
		}
		
	}
}

func createMovie(m http.ResponseWriter , r *http.Request){
	m.Header().Set("content-type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID=strconv.Itoa(rand.Intn(100000))
	movies=append(movies , movie)
	json.NewEncoder(m).Encode(movie)

}

func updateMovie(m http.ResponseWriter , r *http.Request){
	m.Header().Set("content-type", "application/json")
	params:=mux.Vars(r)

	for index, item :=range movies{
		if item.ID==params["id"]{
			movies=append(movies[:index] , movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID=params["id"]
			movies=append(movies , movie)
			json.NewEncoder(m).Encode(movie)
		}
	}

}

func main(){
	r:=mux.NewRouter()

	movies = append(movies, Movie{ID: "01", Language: "Hindi", Name: "Bagwan", Director: &Director{FirstName: "Pritam", LastName: "Kumari"}})
	movies=append(movies, Movie{ID: "02", Language: "marathi", Name: "Jhodha Akbar", Director: &Director{FirstName: "Sweety", LastName: "Singh"}})
	r.HandleFunc("/movies" , getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}" , getMovieByID).Methods("GET")
	r.HandleFunc("/movies/{id}" , updateMovie).Methods("PUT")
	r.HandleFunc("/movie" , createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", DeleteMovie).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080" , r))
}