package main

import (
    "net/http"
    "os"
    "bytes"
    "time"
)

func someHandler(w http.ResponseWriter, r *http.Request){
    // [Vulnerability] Hardcoded password
    password := "geheim"
    sysInfo, _ := os.Executable()

    var c http.Client

    // [Vulnerability] Privacy Violation
    _, err := c.Get(password)
    if err != nil { panic(err.Error()) }

    // [Vulnerability] System Information Leak: External
    _, err = c.Head(sysInfo)
    if err != nil { panic(err.Error()) }

    // [Vulnerability] Privacy Violation
    _, err = c.Post("http://ctf.pwntester.com", "stuff/morestuff", bytes.NewBufferString(password))
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
