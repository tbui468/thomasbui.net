package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
)

type PageData struct {
    Title   string
    Heading string
    Message string
}

func handler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("/home/thomas/blog/templates/index.html")
    if err != nil {
        log.Fatal(err)        
    }

    data := PageData{
        Title: "Blog About Stuff",
        Heading: "Welcome!",
        Message: "test",
    }
    tmpl.Execute(w, data)
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Server listening on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println(err)
    }
}
