package viewmodel

import (
    "github.com/bfraz/webapp/model"
)

type BeePage struct {
    Title string
    Bees []Bee
    Active string
}

type Bee struct {
    Name string
    ImageURL string
    Description string
    IsOrientRight bool
}

func NewBee(b []model.Bee) BeePage {
    result := BeePage{
        Title: "Humble Homesteaders - Bee",
        Active: "bee",
    }
    result.Bees = make([]Bee, len(b))
    for i := 0; i < len(b); i++ {
        vm := beetoVM(b[i])
        vm.IsOrientRight = i%2 == 1
        result.Bees[i] = vm
    }
    return result
}

func beetoVM(b model.Bee) Bee {
    return Bee {
        Name: b.Name,
        ImageURL: b.ImageURL,
        Description: b.Description,
    }
}
