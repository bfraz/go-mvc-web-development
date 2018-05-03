package controller

import (
	"html/template"
	"net/http"
)

var (
	homeController      home
	livestockController livestock
	beeController	    bee
	fruitvegController  fruitveg
)

func Startup(templates map[string]*template.Template) {
	homeController.homeTemplate = templates["home.html"]
	livestockController.livestockTemplate = templates["livestock.html"]
	livestockController.livestockProductTemplate = templates["livestock_product_detail.html"]
	beeController.beeTemplate = templates["bee.html"]
	fruitvegController.fruitvegTemplate = templates["fruitveg.html"]
	homeController.registerRoutes()
	livestockController.registerRoutes()
	beeController.registerRoutes()
	fruitvegController.registerRoutes()
	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
}
