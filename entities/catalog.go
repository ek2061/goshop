package entities

import "fmt"

type Catalog struct {
	Catalog_Id int64
	Name       string
	Hiden      uint
}

func (catalog Catalog) ToString() string {
	return fmt.Sprintf("cid: %d\nname: %s\nhiden: %d\n",
		catalog.Catalog_Id,
		catalog.Name,
		catalog.Hiden)
}
