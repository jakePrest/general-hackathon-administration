package main

import (
    "github.com/montanaflynn/stats"
    "net/http"
    "github.com/gorilla/mux"
    "log"
    "strconv"
    "fmt"
    "html/template"
)

 func Index(w http.ResponseWriter, r *http.Request) {
   t := template.New("index.html") 
   t, _ = t.ParseFiles("./index.html")  // Parse template file.
   t.Execute(w, nil)
}

type DataSet struct {
  Data []float64
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
    r.HandleFunc("/", Index)
    r.HandleFunc("/api/process", ProcessData)
    err := http.ListenAndServe(":7000", r)
    if err != nil {
      log.Fatal(err)
    }
}
