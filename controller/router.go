package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
)

func top(w http.ResponseWriter, r *http.Request) {
	// t := template.Must(template.ParseFiles("view/top.html"))
	// str := "Sample Message"
	// if err := t.ExecuteTemplate(w, "top.html", str); err != nil {
	// 	log.Fatal(err)
	// }
	http.Redirect(w, r, "/upload", 301)
}

func logger(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Handler called - %T", h)
		h(w, r)
	})
}

func Router(db *sqlx.DB) {
	zc := newZipController(db)
	http.HandleFunc("/", top)
	http.HandleFunc("/upload", zc.upload)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
