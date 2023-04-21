package main

import (
    "net/http"
    "os/exec"
    "time"
)

func performLs1(a string){
    // [Vulnerability] Command Injection (called from "someHandler" with tainted "a")
    err := exec.Command("ls", a).Run()
    if err != nil { panic(err.Error()) }
}

func performLs2(a string){
    // [Safe] Not command injection, we won't call this function with tainted data
    err := exec.Command("ls", a).Run()
    if err != nil { panic(err.Error()) }
}

func someHandler(w http.ResponseWriter, r *http.Request){
    k1 := r.FormValue("input")

    // [Vulnerability] Command Injection
    err := exec.Command("ls", k1).Run()
    if err != nil { panic(err.Error()) }

    performLs1(k1)
    performLs2(".")
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
