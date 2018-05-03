package controller

import (
  "html/template"
  "net/http"
  "github.com/bfraz/webapp/model"
  "github.com/bfraz/webapp/viewmodel"
)
type fruitveg struct {
  fruitvegTemplate *template.Template
}

func (fv fruitveg) registerRoutes(){
  http.HandleFunc("/fruitveg", fv.handleFruitVeg)
  http.HandleFunc("/fruitveg/", fv.handleFruitVeg)
}

func (fv fruitveg) handleFruitVeg(w http.ResponseWriter, r *http.Request){
  fruitvegs := model.GetFruitVegs()
  vm := viewmodel.NewFruitVeg(fruitvegs)
  w.Header().Add("Content-Type", "text/html")
  fv.fruitvegTemplate.Execute(w,vm)
}
