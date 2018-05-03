package controller

import (
    "html/template"
    "net/http"
    "github.com/bfraz/webapp/model"
    "github.com/bfraz/webapp/viewmodel"
)

type bee struct {
    beeTemplate *template.Template
}

func (b bee) registerRoutes(){
    http.HandleFunc("/bee", b.handleBee)
    http.HandleFunc("/bee/", b.handleBee)
}

func (b bee) handleBee(w http.ResponseWriter, r *http.Request){
    bees := model.GetBees()
    vm := viewmodel.NewBee(bees)
    w.Header().Add("Content-Type", "text/html")
    b.beeTemplate.Execute(w,vm)
}
