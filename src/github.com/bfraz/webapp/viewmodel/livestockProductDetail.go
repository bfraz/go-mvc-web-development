package viewmodel

import (
	"github.com/bfraz/webapp/model"
)

type LivestockProductDetail struct {
	Title    string
	Active   string
	LivestockProducts []LivestockProduct
}
type LivestockProduct struct {
  	Name        string
  	Description string
  	Price       float32
  	ImageURL    string

}

func NewLivestockProductDetail(lsps []model.LivestockProduct) LivestockProductDetail {
	result := LivestockProductDetail{
		Title:    "Humble Homesteaders",
		Active:   "lsproducts",
		LivestockProducts: []LivestockProduct{},
	}
	for _, lsp := range lsps {
		result.LivestockProducts = append(result.LivestockProducts, lsptoVM(lsp))
	}
	return result
}

func lsptoVM(lsp model.LivestockProduct) LivestockProduct {
	return LivestockProduct{
		Name:        lsp.Name,
		ImageURL:    lsp.ImageURL,
        Price:       lsp.Price,
		Description: lsp.Description,
	}
}
