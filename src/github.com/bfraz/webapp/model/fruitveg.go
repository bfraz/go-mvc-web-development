package model


type FruitVeg struct {
	Name             string
	DescriptionShort string
	Price   float32
	ImageURL         string
	ID               int
}

var fruitvegs []FruitVeg = []FruitVeg{FruitVeg{
	Name:             "Lemons",
	DescriptionShort: "farm fresh grown lemons.",
	Price:   1.09,
	ImageURL:        "lemon.png",
	ID:              1,
}, FruitVeg{
	Name:             "Apples",
	DescriptionShort: "Honeycrisp, gala, and more!",
	Price:    0.89,
	ImageURL:         "apple.png",
	ID:               2,
}, FruitVeg{
	Name:             "Watermelons",
	DescriptionShort: "From the plains of Missouri.",
	Price:    3.99,
	ImageURL:         "watermelon.png",
	ID:               3,
}, FruitVeg{
	Name:             "Kiwis",
	DescriptionShort: "Mmmhmm...delicious",
	Price:    5.99,
	ImageURL:         "kiwi.png",
	ID:               4,
}, FruitVeg{
	Name:             "Oranges",
	DescriptionShort: "Nothing rhymes with Orange?",
	Price:    1.67,
	ImageURL:         "orange.png",
	ID:               6,
}, FruitVeg{
	Name:             "Strawberries",
	DescriptionShort: "The perfect balance of sweet and tart.",
	Price:    4.36,
	ImageURL:         "strawberry.png",
	ID:               8,
},
}

func GetFruitVegs() []FruitVeg {
	return fruitvegs
}
