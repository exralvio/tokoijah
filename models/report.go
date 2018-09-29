package models

type ProductItem struct {
	Sku			string		`json:"sku"`
	Name		string		`json:"name"`
	Qty			int			`json:"qty"`
	Average		int		`json:"average"`
	Total		int		`json:"total"`
}

type SaleItem struct {
	Invoice			string		`json:"invoice"`
	Date			string		`json:"date"`
	Sku				string		`json:"sku"`
	Name			string		`json:"name"`
	Qty				int			`json:"qty"`
	SalePrice		float64		`json:"sale_price"`
	Total			float64		`json:"total"`
	BuyPrice		float64		`json:"buy_price"`
	Profit			float64		`json:"profit"`
}

func SumBuyPrice(x[] *Purchase) float64{
	if len(x) == 0 {
		return 0
	}

	var total_price, result float64
	var total_receive int

	for _, value := range x {
		total_price += value.Total
		total_receive += value.NumberReceive
	}

	result = total_price / float64(total_receive)

	return result
}