package main

import (
  "os"
  "log"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request)  {
  name, err := os.Hostname()
  if err != nil {
    panic(err)
  }
  w.Write([]byte(name))
}

func main()  {
  http.HandleFunc("/", handler)
  log.Fatal(http.ListenAndServe(":8089", nil))
}
