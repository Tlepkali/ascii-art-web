package main

import (
	"errors"
	"log"
	"net/http"
	"text/template"

	ascii "ascii-art-web/output"
)

type output struct {
	HasOutput   bool
	Color       string
	InitialText string
	Text        string
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		err := errors.New("404\nPage not found")
		ErrorPage(w, err, 404)
		return
	}

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		err := errors.New("405" + "\n" + "Method not allowed")
		ErrorPage(w, err, 405)
		return
	}

	ts, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		log.Println(err.Error())
		err = errors.New("500" + "\n" + "Internal Server Error")
		ErrorPage(w, err, 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		ErrorPage(w, err, 500)
		return
	}
}

func AsciiDraw(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ascii" {
		err := errors.New("404\nPage not found")
		ErrorPage(w, err, 404)
		return
	}
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		err := errors.New("405" + "\n" + "Method not allowed")
		ErrorPage(w, err, 405)
		return
	}

	text := r.FormValue("text")
	font := r.FormValue("font")
	color := r.FormValue("color")

	textWithFont, err := ascii.Asciiart(text, font)

	t, err1 := template.ParseFiles("./templates/index.html")
	if err1 != nil {
		err1 = errors.New("500" + "\n" + "Internal Server Error")
		ErrorPage(w, err1, 500)
		return
	}

	if err != nil {
		log.Println(err.Error())
		err = errors.New("400" + "\n" + "Bad request")
		ErrorPage(w, err, 400)
		return
	}

	if text == "" || font == "" {
		t.Execute(w, err)
		return
	}

	data := output{true, color, text, textWithFont}

	err = t.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		err = errors.New("500" + "\n" + "Internal Server Error")
		ErrorPage(w, err, 500)
		return
	}
}

func ErrorPage(w http.ResponseWriter, err error, code int) {
	w.WriteHeader(code)
	t, err1 := template.ParseFiles("./templates/error.html")
	if err1 != nil {
		http.Error(w, "500"+"\n"+"Internal Server Error", http.StatusInternalServerError)
		return
	}
	t.Execute(w, err)
	return
}
