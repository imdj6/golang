package main

import "net/http"


func main() {	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Danish Jamal!"))
	})
	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Goodbye, Danish Jamal!"))
	})

	http.ListenAndServe(":8082", nil)
}