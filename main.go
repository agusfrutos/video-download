package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", startHandler)
	http.HandleFunc("/fmt", testHandler)
	fmt.Printf("Listening on port 3000\n")
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}

func startHandler(w http.ResponseWriter, r *http.Request) {
	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	index := "index.html"

	tmpl, err := template.ParseFiles(index)
	check(err)

	// inicio, err := os.ReadFile("index.html")
	check(err)

	// t, err := template.New("inicio").Parse(home)
	// check(err)

	tmpl.Execute(w, nil)

}

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("get realizado")
}
