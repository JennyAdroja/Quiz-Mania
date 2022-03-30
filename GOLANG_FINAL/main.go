package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)

}
func gkPage(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "gk.html", nil)

}
func computerPage(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "computer.html", nil)

}
func sciencePage(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "science.html", nil)

}
func sportsPage(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "sports.html", nil)

}
func geographyPage(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "geography.html", nil)

}
func filmsPage(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "films.html", nil)

}
func main() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))
	http.HandleFunc("/", homePage)
	http.HandleFunc("/gk", gkPage)
	http.HandleFunc("/computer", computerPage)
	http.HandleFunc("/science", sciencePage)
	http.HandleFunc("/sports", sportsPage)
	http.HandleFunc("/geography", geographyPage)
	http.HandleFunc("/films", filmsPage)
	log.Println("Starting the web server at port 8888")
	http.ListenAndServe(":8888", nil)
}
