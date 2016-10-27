package main

import (
	"net/http"
	"io"
	"io/ioutil"
	"fmt"
	"encoding/json"
)

func calcSantas(w http.ResponseWriter, r *http.Request) {
	var santa []Santa
	
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &santa); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	fmt.Println(santa)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./web")))
	http.HandleFunc("/santa/", calcSantas)
        http.ListenAndServe(":3000", nil)
}
