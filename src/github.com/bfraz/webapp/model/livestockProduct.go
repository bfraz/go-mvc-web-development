package model


type LivestockProduct struct {
	Name             string
	Description string
	Price            float32
	ImageURL         string
	ID               int
	LivestockID      int
}



func GetLivestockProductsWithLivestockID(livestockID int) ([]LivestockProduct, error) {
	result := []LivestockProduct{}
	rows, err := db.Query(`
		SELECT name, description, price, imageurl
		FROM livestockproduct
		WHERE livestockid = $1`, livestockID)
	if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
    ls := LivestockProduct{}
    if err := rows.Scan(&ls.Name, &ls.Description, &ls.Price, &ls.ImageURL); err != nil {
        return nil, err
    }
        result = append(result, ls)
    }
    return result, nil
}

////////////////////////
// func GetLivestockProductsWithLivestockID(livestockID int) []LivestockProduct {
// 	result := []LivestockProduct{}
//   lps := GetLivestockProducts()
// 	for _, lp := range lps {
// 		if lp.LivestockID == livestockID {
// 			result = append(result, lp)
// 		}
// 	}
// 	return result
// }
//
// func GetLivestockProducts() []LivestockProduct {
//   return livestockProducts
// }
//
// var livestockProducts []LivestockProduct = []LivestockProduct{LivestockProduct{
// 	Name:             "Chick",
// 	Description: "These chicks are born to be wild.",
// 	Price:   0.99,
// 	ImageURL:        "star_chicks.png",
// 	ID:              1,
// 	LivestockID:      3,
// }, LivestockProduct{
// 	Name:             "Eggs/Dozen",
// 	Description: "Yolks so orange, they appear as the setting sun.",
// 	Price:            3,
// 	ImageURL:         "farm_fresh_eggs.png",
// 	ID:               2,
// 	LivestockID:      3,
// }, LivestockProduct{
// 	Name:             "Processed Chicken",
// 	Description: "You won't have problems catching a chicken that's been processed!",
// 	Price:             5,
// 	ImageURL:         "processed_chicken.png",
// 	ID:               3,
// 	LivestockID:       3,
//   },
// }
