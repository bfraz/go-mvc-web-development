package model

type Livestock struct {
    ID          int
    ImageURL    string
    Name        string
    Description string
}

var livestock []Livestock = []Livestock{Livestock{
    ID: 1,
    ImageURL: "sheep.png",
    Name: "Sheep",
    Description: `Let there be wool! Whether you want live sheep to add to your flock
                                    or lamb chops with your parsley.`,
    }, Livestock{
    ID: 2,
    ImageURL: "pigs.png",
    Name: "Pigs/Hogs",
    Description: `That's some pig! We believe a happy hog is healthy hog. Pick up a hog or pig of
                                    your choice, or send them to a select butcher for prime cuts.`,
    }, Livestock{
    ID: 3,
    ImageURL: "chicks.png",
    Name: "Chicks/Chickens",
    Description: `From eggs, to chicks, to whole chickens: you ask, we provide.`,
    },
}

func GetLivestock() []Livestock {
    return livestock
}

// NOTE: This is if we were pulling livestock from the db.
// func GetLivestock() ([]Livestock, error) {
// 	result := []Livestock{}
//   rows, err := db.Query(`SELECT * from livestock`)
//   if err != nil {
//     return nil, err
//   }
//   defer rows.Close()
//   for rows.Next() {
//     ls := Livestock{}
//     if err := rows.Scan(&ls.ID, &ls.ImageURL, &ls.Name, &ls.Description); err != nil {
//       return nil, err
//     }
//     result = append(result, ls)
//   }
// 	return result, nil
// }
