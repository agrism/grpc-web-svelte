package main

import (
	"encoding/json"
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"log"
	"net/http"
)

type Category struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

func main()  {

	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":9003", nil))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func homePage(w http.ResponseWriter, r *http.Request){

	enableCors(&w)

	w.Header().Add("Content-Type", "application/json")

	var cats []Category

	for i:= 0; i < 1000000; i++ {
		cats = append(cats, Category{
			Id: fmt.Sprintf("%v", i),
			Name: randomdata.City(),
		})
	}

	json.NewEncoder(w).Encode(cats)
}