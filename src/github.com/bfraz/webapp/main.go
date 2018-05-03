package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"github.com/bfraz/webapp/controller"
	"github.com/bfraz/webapp/middleware"
	"fmt"
	"log"
	"github.com/bfraz/webapp/model"
	_ "github.com/lib/pq"
	"database/sql"
)

func main() {
	templates := populateTemplates()
	db := connectToDatabase()
	defer db.Close() //close db connection gracefully when program exits
	controller.Startup(templates)
	http.ListenAndServeTLS(":8000", "cert.pem", "key.pem", &middleware.TimeoutMiddleware{new(middleware.GzipMiddleware)})
}

func connectToDatabase() *sql.DB {
	// username is postgres, password is admin, and db is named bfraz.
	postgresConnectionString := "postgres://postgres:admin@localhost/bfraz?sslmode=disable"
	db, err := sql.Open("postgres", postgresConnectionString)
	if err != nil {
		log.Fatalln(fmt.Errorf("Unable to connect to database: %v", err))
	}
	model.SetDatabase(db)
	return db
}

func populateTemplates() map[string]*template.Template {
	result := make(map[string]*template.Template)
	const basePath = "templates"
	layout := template.Must(template.ParseFiles(basePath + "/_layout.html"))
	template.Must(layout.ParseFiles(basePath+"/_header.html"))
	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error())
	}
	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directory: " + err.Error())
	}
	for _, fi := range fis {
		f, err := os.Open(basePath + "/content/" + fi.Name())
		if err != nil {
			panic("Failed to open template '" + fi.Name() + "'")
		}
		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Failed to read content from file '" + fi.Name() + "'")
		}
		f.Close()
		tmpl := template.Must(layout.Clone())
		_, err = tmpl.Parse(string(content))
		if err != nil {
			panic("Failed to parse contents of '" + fi.Name() + "' as template")
		}
		result[fi.Name()] = tmpl
	}
	return result
}
