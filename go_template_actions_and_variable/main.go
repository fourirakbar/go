package main

import "fmt"
import "html/template"
import "net/http"

type info struct {
	Affiliation string
	Address     string
}

type Person struct {
	Name    string
	Gender  string
	Hobbies []string
	Info    info
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var person = Person{
			Name:    "Fourir Akbar",
			Gender:  "Male",
			Hobbies: []string{"String A, String B, String C"},
			Info:    info{"Info Satu", "Info yg kedua"},
		}

		var tmpl = template.Must(template.ParseFiles("view.html"))
		if err := tmpl.Execute(w, person); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("server started at localhost 9000")
	http.ListenAndServe(":9000", nil)
}

func (t info) GetAffiliationDetailInfo() string {
	return "Have 31 divisions"
}
