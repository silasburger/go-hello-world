package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type favoritePerson struct {
	Name          string
	FavoriteColor string
	Nickname      string
}

type favoritePeople []favoritePerson

func helloPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hompage endpoint hit")
}

func allFavoritePeople(w http.ResponseWriter, r *http.Request) {
	favorites := favoritePeople{
		favoritePerson{Name: "Kona", FavoriteColor: "brown", Nickname: "Ko"},
		favoritePerson{Name: "Neville", FavoriteColor: "red", Nickname: "Nev"},
	}
	tmpl, err := template.ParseFiles("layout.html")
	if err != nil {
		fmt.Println(err)
	}
	tmpl.Execute(w, favorites)

}

func handleRequests() {
	http.HandleFunc("/", helloPage)
	http.HandleFunc("/friends", allFavoritePeople)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func main() {
	handleRequests()
}
