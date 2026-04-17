package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"  
    "strconv"
)


func home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }

    files := []string{
        "./ui/html/base.html",
        "./ui/html/partials/nav.html",
        "./ui/html/pages/home.html",
    }

    //parse file into template sets
    ts, err := template.ParseFiles(files...)
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error", 500)
        return
    }

    //use Execute to write html into response body
    //use ExecuteTemplate to write template into response body
    err = ts.ExecuteTemplate(w, "base", nil)
    if err != nil {
        log.Println(err.Error())
        http.Error(w, "Internal Server Error", 500)
    }
}

func question(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Query().Get("id"))
    if err != nil || id < 1 {
        http.NotFound(w, r)
        return
    }

    fmt.Fprintf(w, "Display a specific question with ID %d...", id)
}
