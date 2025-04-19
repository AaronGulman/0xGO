package main

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"log"
)

func main(){
	r :=chi.NewRouter()

	r.Get("/govsnode",func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("I think this might be GO time"))
	})

	if err := http.ListenAndServe(":8080",r); err!=nil {
		log.Fatal(err)
	}
}