package controller

import (
    "html/template"
    "net/http"
    "github.com/bfraz/webapp/model"
    "github.com/bfraz/webapp/viewmodel"
    "regexp"
    "strconv"
)

type livestock struct {
    livestockTemplate *template.Template
    livestockProductTemplate *template.Template
}

func (l livestock) registerRoutes() {
    http.HandleFunc("/livestock", l.handleLivestock)
    http.HandleFunc("/livestock/", l.handleLivestock)
}

func (l livestock) handleLivestock(w http.ResponseWriter, r *http.Request){
    livestockProductPattern, _ := regexp.Compile(`/livestock/(\d+)`)
    matches := livestockProductPattern.FindStringSubmatch(r.URL.Path)
    executeLivestockVM := false
    if len(matches) > 0 {
        livestockProductID, _ := strconv.Atoi(matches[1])
        lsps, err := model.GetLivestockProductsWithLivestockID(livestockProductID)
        if err != nil {
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        }
        if len(lsps) > 0 {
            l.handleLivestockProduct(w, r, livestockProductID)
        } else{
            executeLivestockVM = true
        }
    } else{
        executeLivestockVM = true
    }

    if (executeLivestockVM) {
        livestock := model.GetLivestock()
        vm := viewmodel.NewLivestock(livestock)
        w.Header().Add("Content-Type", "text/html")
        l.livestockTemplate.Execute(w,vm)
    }
}


func (l livestock) handleLivestockProduct(w http.ResponseWriter, r *http.Request, livestockProductID int) {
	lsps, err := model.GetLivestockProductsWithLivestockID(livestockProductID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
	vm := viewmodel.NewLivestockProductDetail(lsps)
	w.Header().Add("Content-Type", "text/html")
	l.livestockProductTemplate.Execute(w, vm)
}
