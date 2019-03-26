package main

import(
    "fmt"
    "net/http"
)


func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":9090", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello World!")
}

