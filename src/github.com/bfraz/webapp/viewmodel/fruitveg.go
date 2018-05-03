package viewmodel

import (
  "github.com/bfraz/webapp/model"
)

type FruitVegPage struct {
  Title    string
  Active   string
  FruitVegs []FruitVeg
}

type FruitVeg struct {
	Name             string
	DescriptionShort string
	Price   float32
	ImageURL         string
}

func NewFruitVeg(fvs []model.FruitVeg) FruitVegPage{
  result := FruitVegPage {
    Title:    "Humble Homesteaders",
    Active:   "fruitveg",
    FruitVegs: []FruitVeg{},
  }
  for _, fv := range fvs {
		result.FruitVegs = append(result.FruitVegs, fvToVM(&fv))
	}
  return result
}

func fvToVM(fv *model.FruitVeg) FruitVeg {
	return FruitVeg{
		Name:             fv.Name,
		DescriptionShort: fv.DescriptionShort,
		Price:            fv.Price,
		ImageURL:         fv.ImageURL,
	}
}
