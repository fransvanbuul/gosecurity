package main

import (
    "net/http"
    "time"
)

func someHandler(w http.ResponseWriter, r *http.Request){
    var c http.Client

    aUrl := r.FormValue("url")
    // [Vulnerability] Server-Side Request Forgery
    _, err := c.Get(aUrl)
    if err != nil { panic(err.Error()) }
    // [Vulnerability] Server-Side Request Forgery
    _, err = c.Post(aUrl, "something/thing", nil)
    if err != nil { panic(err.Error()) }
}

func main(){
    srv := &http.Server{
        ReadTimeout: 5 * time.Second,
        WriteTimeout: 10 * time.Second,
        Handler: http.HandlerFunc(someHandler),
    }
    // [Vulnerability] Insecure transport
    err := srv.ListenAndServe()
    if err != nil { panic(err.Error()) }
}
