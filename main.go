package main

import (
	"fmt"
	"net/http"
    "log"
    //"html/template"
    "io/ioutil"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data, err := ioutil.ReadFile("templates/main.html")
        if err != nil {
            panic("Unable to read templates/main.html")
        }
        fmt.Fprintf(w, "%s", data)
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}
