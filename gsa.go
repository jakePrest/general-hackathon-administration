package main

import (
    //"github.com/montanaflynn/stats"
    "net/http"
    "github.com/gorilla/mux"
    "log"
)

 func Index(w http.ResponseWriter, r *http.Request) {
	     w.Write([]byte("Testing API"))
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/api/test", Index)
    err := http.ListenAndServe(":5000", r)
    if err != nil {
      log.Fatal(err)
    }
}
