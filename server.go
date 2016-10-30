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

const MAXITTERATIONS = 1000

// Http handler to calc secret santas from json
func calcSantas(w http.ResponseWriter, r *http.Request) {
	var santas []Santa
	itterations := 1
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
	} else {
		// Generate secret santas
		for (!generateSantas(santas) && itterations < MAXITTERATIONS) {
			itterations++
			for i := range santas {
				santas[i].Selectable = nil
				santas[i].Selected = santas[i].Id
			}
		}

		if (itterations == MAXITTERATIONS) {
			w.WriteHeader(417)
		} else {
			w.WriteHeader(200)
			printSantas(santas)
		}
	}
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

// Send sms with 49elks API
func sendSMS(from Santa, to Santa) {
	data := url.Values{
        "from": {"Tomteverkstan"},
        "to": {from.Phone},
        "message":{fmt.Sprintf("Hej %s, du ska köpa ett paket till %s. God jul!", from.Name, to.Name)}}

    req, err := http.NewRequest("POST", "https://api.46elks.com/a1/SMS", bytes.NewBufferString(data.Encode()))
    req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
    req.SetBasicAuth("<API Username>", "<API Password>")

    client := &http.Client{}
    resp, err := client.Do(req)

    defer resp.Body.Close()
    _, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        fmt.Println("Oh dear!!!")
    }
}

// Main function
func main() {
	http.Handle("/", http.FileServer(http.Dir("./web")))
	http.HandleFunc("/santa/", calcSantas)
        http.ListenAndServe(":3000", nil)
}
