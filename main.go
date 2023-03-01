package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/matheusjean28/go_Restapi/domain"
)

func main() {
	http.HandleFunc("/person/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {

			var person domain.Person
			err := json.NewDecoder(r.Body).Decode(&person)
			if err != nil {
				fmt.Printf("error trying to decode r.body. body must be in a json format %s", err.Error())
				http.Error(w, "error on create a new persn", http.StatusBadRequest)
				return
			}
			if person.ID <= 0 {
				http.Error(w, "Id must be more than zero ", http.StatusBadRequest)
				return
			}

			w.WriteHeader(http.StatusCreated)
			return
		}
		http.Error(w, "not implemented ", http.StatusInternalServerError)
	})

	http.ListenAndServe(":8080", nil)
}
