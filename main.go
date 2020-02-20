package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)
type Artist struct {
	ID           int64    `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int64    `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    Locations
	Dates        Dates
	Relations    Relations
}
type Locations struct {
	Index []Location `json:"index"`
}
type Location struct {
	ID        int64    `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}
type Dates struct {
	Index []Date `json:"index"`
}
type Date struct {
	ID    int64    `json:"id"`
	Dates []string `json:"dates"`
}
type Relations struct {
	Index []Relation `json:"index"`
}
type Relation struct {
	ID             int64                  `json:"id"`
	DatesLocations map[string]interface{} `json:"datesLocations"`
}



var artists []Artist
var locations Locations
var dates Dates
var relations Relations

func main() {
	GetAndUnmarshal("https://groupietrackers.herokuapp.com/api/artists", &artists)
	GetAndUnmarshal("https://groupietrackers.herokuapp.com/api/locations", &locations)
	GetAndUnmarshal("https://groupietrackers.herokuapp.com/api/dates", &dates)
	GetAndUnmarshal("https://groupietrackers.herokuapp.com/api/relation", &relations)
	AppendToStruct()

	fs := http.FileServer(http.Dir("assets"))
	js := http.FileServer(http.Dir("scripts"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", js))

	http.HandleFunc("/", Server)
	http.HandleFunc("/answer", Answer)
	http.HandleFunc("/loc", LocationsArt)
	http.HandleFunc("/dat", DatesArt)
	http.HandleFunc("/rel", RelationsArt)

	port := os.Getenv("PORT")
	if port == "" {
		port = "1998"
	}
	port = ":" + port

	http.ListenAndServe(port, nil)

}
func Server(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		tml, err := template.ParseFiles("err404.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		tml.Execute(w, r.URL.Path[1:])
		return
	}

	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, r.URL.Path[1:])
}

func Answer(w http.ResponseWriter, r *http.Request) {
	art, err := json.Marshal(artists)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(art)
}

func LocationsArt(w http.ResponseWriter, r *http.Request) {
	loc, err := json.Marshal(locations)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(loc)
}
func DatesArt(w http.ResponseWriter, r *http.Request) {
	date, err := json.Marshal(dates)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(date)
}
func RelationsArt(w http.ResponseWriter, r *http.Request) {
	rel, err := json.Marshal(relations)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(rel)
}
func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}
func GetAndUnmarshal(api string, arr interface{}) {
	APIurl := api
	req, _ := http.NewRequest("GET", APIurl, nil)

	res, _ := http.DefaultClient.Do(req)
	fmt.Println(res)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	json.Unmarshal([]byte(body), &arr)
}
func AppendToStruct() {
	for index := range locations.Index {

		artists[index].Locations.Index = append(artists[index].Locations.Index, locations.Index[index])
	}
	for index := range dates.Index {

		artists[index].Dates.Index = append(artists[index].Dates.Index, dates.Index[index])
	}
	for index := range relations.Index {

		artists[index].Relations.Index = append(artists[index].Relations.Index, relations.Index[index])
	}
}
