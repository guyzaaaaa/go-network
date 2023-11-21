package main

import (
	"fmt"
	//"io/fs"
	"net/http"
	"github.com/gorilla/mux"
)
func main() {

	r:=mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request){
		vars := mux.Vars(r)
        title := vars["title"]
        page := vars["page"]
		
		b := []byte("<h1>Book:" + title + "</h1><br><h2>: " + page + "</h2>")
	w.Write(b)
        fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hello World")
	})
	
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Thanapat")
	})

fs:=http.FileServer(http.Dir("static/"))
http.Handle("/static/", http.StripPrefix("/static/", fs))
	
http.ListenAndServe(":8080", r)
}