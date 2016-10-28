package main

import (
	"net/http"
	"io"
	"io/ioutil"
	"fmt"
	"encoding/json"
)

func calcSantas(w http.ResponseWriter, r *http.Request) {
	var santas []Santa
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &santas); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	// Fill seletable santas
	for i := range santas {
		santas[i].Selectable = selectableSantas(santas[i], santas)
	}

	fmt.Println(santas)
}

func selectableSantas(santa Santa, santas []Santa) []int {
	selectable := make([]int, 0)
	for i := range santas {
		if !excludeSanta(santa, santas[i]) {
			selectable = append(selectable, santas[i].Id)
		}
	}
	return selectable
}

func excludeSanta(self Santa, other Santa) bool {
	for i := range self.Exclude {
		if self.Exclude[i] == other.Id || self.Id == other.Id {
			return true
		}
	}
	return false
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./web")))
	http.HandleFunc("/santa/", calcSantas)
        http.ListenAndServe(":3000", nil)
}
