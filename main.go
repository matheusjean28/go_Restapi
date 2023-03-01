package main

import "net/http"

func main() {
	http.HandleFunc("/person/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			w.Write([]s)
			println("route post")
		} else if r.Method == "GET" {
			println("method Get")
		}
		http.Error(w, "not implemented ", http.StatusInternalServerError)
	})

	http.ListenAndServe(":8080", nil)
}
