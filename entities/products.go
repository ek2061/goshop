package entities

import "fmt"

type Product struct {
	Product_Id   int64
	Catalog_Id   int64
	Name         string
	Cost         int16
	Price        int16
	Description  string
	On_Sale      uint
	Start_Sell   string
	End_Sell     string
	Catalog_Name string
}

func (product Product) ToString() string {
	return fmt.Sprintf("pid: %d\ncid: %d\nname: %s\ncost: %d\nprice: %d\ndescription: %s\n",
		product.Product_Id,
		product.Catalog_Id,
		product.Name,
		product.Cost,
		product.Price,
		product.Description)
}
