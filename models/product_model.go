package models

import (
	"database/sql"

	"github.com/ek2061/goshop/entities"
)

type ProductModel struct {
	Db *sql.DB
}

// 找出全部的商品
func (productModel ProductModel) FindAll() (Product []entities.Product, err error) {
	// rows, err := productModel.Db.Query("select * from products")
	rows, err := productModel.Db.Query("select products.*, catalogs.name cname from products inner join catalogs on products.catalog_id=catalogs.catalog_id and products.on_sale=1 and CURRENT_TIME BETWEEN start_sell AND end_sell")
	if err != nil {
		return nil, err
	} else {
		var products []entities.Product
		for rows.Next() {
			var product_id int64
			var catalog_id int64
			var name string
			var cost int16
			var price int16
			var description string
			var on_sale uint
			var start_sell string
			var end_sell string
			var catalog_name string
			err2 := rows.Scan(&product_id, &catalog_id, &name, &cost, &price, &description, &on_sale, &start_sell, &end_sell, &catalog_name)
			if err2 != nil {
				return nil, err2
			} else {
				product := entities.Product{
					Product_Id:   product_id,
					Catalog_Id:   catalog_id,
					Name:         name,
					Cost:         cost,
					Price:        price,
					Description:  description,
					On_Sale:      on_sale,
					Start_Sell:   start_sell,
					End_Sell:     end_sell,
					Catalog_Name: catalog_name,
				}
				products = append(products, product)
			}
		}
		return products, nil
	}
}

// 搜尋商品編號
func (productModel ProductModel) Search(id int64) (Product []entities.Product, err error) {
	// rows, err := productModel.Db.Query("select * from products where product_id = ?", id)
	rows, err := productModel.Db.Query("select products.*, catalogs.name cname from products inner join catalogs on products.catalog_id=catalogs.catalog_id and products.product_id=? and products.on_sale=1 and CURRENT_TIME BETWEEN start_sell AND end_sell", id)
	if err != nil {
		return nil, err
	} else {
		var products []entities.Product
		for rows.Next() {
			var product_id int64
			var catalog_id int64
			var name string
			var cost int16
			var price int16
			var description string
			var on_sale uint
			var start_sell string
			var end_sell string
			var catalog_name string
			err2 := rows.Scan(&product_id, &catalog_id, &name, &cost, &price, &description, &on_sale, &start_sell, &end_sell, &catalog_name)
			if err2 != nil {
				return nil, err2
			} else {
				product := entities.Product{
					Product_Id:   product_id,
					Catalog_Id:   catalog_id,
					Name:         name,
					Cost:         cost,
					Price:        price,
					Description:  description,
					On_Sale:      on_sale,
					Start_Sell:   start_sell,
					End_Sell:     end_sell,
					Catalog_Name: catalog_name,
				}
				products = append(products, product)
			}
		}
		return products, nil
	}
}

// 建立商品資料
func (productModel ProductModel) Create(Product *entities.Product) (err error) {
	result, err := productModel.Db.Exec("insert into products(catalog_id, name, cost, price, description, on_sale, start_sell, end_sell) values(?, ?, ?, ?, ?, ?, ?, ?)",
		Product.Catalog_Id, Product.Name, Product.Cost, Product.Price, Product.Description, Product.On_Sale, Product.Start_Sell, Product.End_Sell)
	if err != nil {
		return err
	} else {
		Product.Product_Id, _ = result.LastInsertId()
		return nil
	}
}

// 更新商品資料
func (productModel ProductModel) Update(Product *entities.Product) (int64, error) {
	result, err := productModel.Db.Exec("update products set catalog_id=?, name=?, cost=?, price=?, description=?, on_sale=?, start_sell=?, end_sell=? where product_id=?",
		Product.Catalog_Id, Product.Name, Product.Cost, Product.Price, Product.Description, Product.On_Sale, Product.Start_Sell, Product.End_Sell, Product.Product_Id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}

func (productModel ProductModel) Delete(id int64) (int64, error) {
	result, err := productModel.Db.Exec("delete from products where product_id = ?", id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}
