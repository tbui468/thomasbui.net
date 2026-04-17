package main

import (
    "log"
    "net/http"
)


func main() {
    mux := http.NewServeMux() //servemux == router
    fileServer := http.FileServer(http.Dir("./ui/static/"))
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))
    mux.HandleFunc("/", home)
    mux.HandleFunc("/question", question)
    log.Println("Starting server on :4000")
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}
