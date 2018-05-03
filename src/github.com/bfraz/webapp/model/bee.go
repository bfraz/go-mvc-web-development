package model

type Bee struct {
    ID          int
    ImageURL    string
    Name        string
    Description string
}

var bees []Bee = []Bee{
    Bee{
        ImageURL: "honey.png",
        Name: "Honey",
        Description: `'A dolla makes me holla honey boo boo'.
                      Get your honey and holla too!`,
    },
    Bee{
        ImageURL: "marked_queen_bee.png",
        Name: "Marked Queen Bee",
        Description: `She's a queen and she demands to be treated like a queen.
                      Start your colony with a true queen! `,

    },
    Bee{
        ImageURL: "bee_hives.png",
        Name: "Bee Hives",
        Description: `Be a stileto in a room full of flats.
                      Let your Bees 'bee' in-style.`,

    },

}

func GetBees() []Bee {
    return bees
}
