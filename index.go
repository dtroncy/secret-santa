package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Participant struct {
	Name  string
	Email string
	Draw  string
}

// index is the handler responsible for rending the index page for the site.
func index(w http.ResponseWriter, r *http.Request) {

	values := r.URL.Query()

	if len(values) != 0 {

		participants := make([]Participant, len(values)/2)

		for k, v := range values {

			position, _ := strconv.Atoi(k[len(k)-1:])

			if strings.Contains(k, "name") {

				participants[position-1].Name = v[0]

			} else if strings.Contains(k, "email") {

				participants[position-1].Email = v[0]

			}

		}

		tirage := make([]string, len(participants))

		for i := 0; i < len(participants); i++ {
			tirage[i] = participants[i].Name
		}

		for i := 0; i < len(participants); i++ {

			if len(tirage) > 1 {

				rand.Seed(time.Now().UnixNano())

				n := rand.Intn(len(tirage)-1) + 1

				participants[i].Draw = tirage[n]

				if n != len(tirage)-1 {

					// Remove the element at index n from tirage.
					tirage[n] = tirage[len(tirage)-1] // Copy last element to index n.

				}

				tirage = tirage[:len(tirage)-1] // Truncate slice.

				if participants[i+1].Name != tirage[0] {

					for j := 1; j < len(tirage); j++ {

						if participants[i+1].Name == tirage[j] {

							t := tirage[0]

							tirage[0] = tirage[j]

							tirage[j] = t

							break
						}
					}

				}

			} else {

				participants[i].Draw = tirage[0]

			}

		}
		fmt.Println("---- Draw result ----")

		for i := 0; i < len(participants); i++ {

			fmt.Println("Participant : " + participants[i].Name + " " + participants[i].Email + " --> " + participants[i].Draw)

			//TODO - sending email

		}

		fmt.Println("-------------------------")
	}
	b := struct {
		Title template.HTML
	}{
		Title: template.HTML("Secret Santa"),
	}

	var templates = template.Must(template.ParseFiles("templates/index.gohtml"))

	err := templates.ExecuteTemplate(w, "index", &b)
	if err != nil {
		http.Error(w, fmt.Sprintf("index : Parsing du template impossible : %v", err), http.StatusInternalServerError)
		return
	}

}
