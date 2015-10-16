package main

import (
    "github.com/montanaflynn/stats"
    "net/http"
    "github.com/gorilla/mux"
    "log"
    "strconv"
    "fmt"
)

 func Index(w http.ResponseWriter, r *http.Request) {
	     w.Write([]byte("Hello World"))
}

/* Placeholder for data processing function */
func ProcessData(w http.ResponseWriter, r *http.Request) {
  var data = []float64{1, 2, 3, 4, 4, 5}
  median, err := stats.Median(data)
  if err != nil {
    log.Println(err)
  }
  medianStr := strconv.FormatFloat(median, 'f', 5, 64)
  fmt.Println(medianStr)
  w.Write([]byte(medianStr))
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/api/test", ProcessData)
    r.HandleFunc("/api/process", ProcessData)
    err := http.ListenAndServe(":5000", r)
    if err != nil {
      log.Fatal(err)
    }
}
