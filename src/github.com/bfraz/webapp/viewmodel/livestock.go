package viewmodel

import (
  "github.com/bfraz/webapp/model"
  "fmt"
)

type LivestockPage struct {
  Title   string
  Livestock []Livestock
  Active string
}

type Livestock struct {
  Name string
  ImageURL string
  Description string
  IsOrientRight bool
  URL string
}

func NewLivestock(livestock []model.Livestock) LivestockPage {
  result := LivestockPage{
    Title:  "Humble Homesteaders - Livestock",
    Active: "livestock",
  }
  result.Livestock = make([]Livestock, len(livestock))
  for i :=0; i < len(livestock); i++ {
    vm := livestocktoVM(livestock[i])
    vm.IsOrientRight = i%2 == 0
    result.Livestock[i] = vm
  }
  return result
}

func livestocktoVM(l model.Livestock) Livestock {
	return Livestock{
    URL:         fmt.Sprintf("/livestock/%v", l.ID),
		Name:        l.Name,
		ImageURL:    l.ImageURL,
		Description:       l.Description,
	}
}
