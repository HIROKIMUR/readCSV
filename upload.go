package main

import (
	// "io/ioutil"
	"encoding/csv"
	"html/template"
	"log"
	"net/http"

	// "os"
	// "path/filepath"
	"fmt"
)

type layout struct {
	layaut1 string
	layaut2 string
	layaut3 string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("html/upload.html"))
}

func uploadHandler(w http.ResponseWriter, req *http.Request) {
	file, _, err := req.FormFile("csvFile")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var line []string
	layauts := make([]layout, 0)
	for {
		i := 0
		line, err = reader.Read()
		if err != nil {
			break
		}
		layauts = append(layauts, layout{line[0], line[1], line[2]})
		i++
	}
	fmt.Println(layauts[1].layaut1)

	// line, err = reader.ReadAll()
	// fmt.Println(line)

	err2 := tpl.ExecuteTemplate(w, "upload.html", layauts)
	if err2 != nil {
		log.Fatalln(err2)
	}

	w.Header()["Location"] = []string{"/upload"}
	w.WriteHeader(http.StatusTemporaryRedirect)
}
