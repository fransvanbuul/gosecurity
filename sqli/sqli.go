package main

import (
    "net/http"
    "database/sql"
    "time"
)

func someHandler(w http.ResponseWriter, r *http.Request){
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/hello")
    if err != nil { panic(err.Error()) }
	defer db.Close()

    name := r.FormValue("name")
    statement := "INSERT INTO users(name) VALUES(" + name + ")"
    // [Vulnerability] SQL injection
    prep, err := db.Prepare(statement)
    if err != nil { panic(err.Error()) }
    _, err = prep.Exec()
    if err != nil { panic(err.Error()) }

    name = "x"
    // [Safe] No SQL injection, "name" cannot be influenced by a user
    prep2, err := db.Prepare("INSERT INTO users(name) VALUES(" + name + ")")
    if err != nil { panic(err.Error()) }
    _, err = prep2.Exec()
    if err != nil { panic(err.Error()) }
}

func main(){
    srv := &http.Server{
        ReadTimeout: 5 * time.Second,
        WriteTimeout: 10 * time.Second,
        Handler: http.HandlerFunc(someHandler),
    }
    err := srv.ListenAndServe()
    if err != nil { panic(err.Error()) }
}
