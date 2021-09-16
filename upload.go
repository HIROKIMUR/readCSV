package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type layout struct {
	Layout1 string
	Layout2 string
	Layout3 string
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
	// layauts := make([]layout, 0)
	var layouts []layout
	for {
		i := 0
		line, err = reader.Read()
		if err != nil {
			break
		}
		layouts = append(layouts, layout{line[0], line[1], line[2]})
		i++
	}
	fmt.Println(layouts[0].Layout1)

	err2 := tpl.ExecuteTemplate(w, "upload.html", layouts)
	if err2 != nil {
		log.Fatalln(err2)
	}

	// w.Header()["Location"] = []string{"/upload"}
	// w.WriteHeader(http.StatusTemporaryRedirect)
}
