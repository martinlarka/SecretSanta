package main

import (
        "net/http"
	"fmt"
)

func calcSantas(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./web")))
	http.HandleFunc("/santa/", calcSantas)
        http.ListenAndServe(":3000", nil)
}
