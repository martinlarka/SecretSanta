package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

// Http handler to calc secret santas from json
func calcSantas(w http.ResponseWriter, r *http.Request) {
	var santas []Santa
	itterations := 0
	rand.Seed(time.Now().UTC().UnixNano())
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
	// Generate secret santas
	for (!generateSantas(santas) && itterations < 50) {
		itterations++
		for i := range santas {
			santas[i].Selectable = nil
			santas[i].Selected = santas[i].Id
		}
	}
	fmt.Printf("itterations: %d\n", itterations)
	printSantas(santas)		
}

// Generate selected santas for each santa
func generateSantas(santas []Santa) bool {
	// Fill selectable santas
	
	var selectedSantas []int
	selectOrder := rand.Perm(len(santas))
	for _, v := range selectOrder {
		// Pick random santa
		santas[v].Selectable = selectableSantas(santas[v], santas, selectedSantas)
		if len(santas[v].Selectable) == 0 {
			return false
		}
		s := rand.Intn(len(santas[v].Selectable))
		santas[v].Selected = santas[v].Selectable[s]
		selectedSantas = append(selectedSantas, santas[v].Selected)
	}
	return true
}

// Populate selectable santa array
func selectableSantas(santa Santa, santas []Santa, selectedSantas []int) []int {
	selectable := make([]int, 0)
	for i := range santas {
		if !excludedSanta(santa, santas[i]) && !intInArray(santas[i].Id, selectedSantas) && santa.Id != santas[i].Id {
			selectable = append(selectable, santas[i].Id)
		}
	}
	return selectable
}

// Evaluates if santa should be excluded
func excludedSanta(self Santa, other Santa) bool {
	for i := range self.Exclude {
		if self.Exclude[i] == other.Id || self.Id == other.Id {
			return true
		}
	}
	return false
}

// Evaluates if int is in array. Used to find santa id in array of already selected santas
func intInArray(a int, list []int) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

// Print result for testing
func printSantas(santas []Santa) {
	for i := range santas {
			fmt.Printf("%s gives to %s \n", santas[i].Name, santas[santas[i].Selected].Name)
		}
}

// Main function
func main() {
	http.Handle("/", http.FileServer(http.Dir("./web")))
	http.HandleFunc("/santa/", calcSantas)
        http.ListenAndServe(":3000", nil)
}
