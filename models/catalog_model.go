package models

import (
	"database/sql"

	"github.com/ek2061/goshop/entities"
)

type CatalogModel struct {
	Db *sql.DB
}

// 找出全部的目錄
func (catalogModel CatalogModel) FindAll() (Catalog []entities.Catalog, err error) {
	rows, err := catalogModel.Db.Query("select * from catalogs where hiden=0")
	if err != nil {
		return nil, err
	} else {
		var catalogs []entities.Catalog
		for rows.Next() {
			var catalog_id int64
			var name string
			var hiden uint
			err2 := rows.Scan(&catalog_id, &name, &hiden)
			if err2 != nil {
				return nil, err2
			} else {
				catalog := entities.Catalog{
					Catalog_Id: catalog_id,
					Name:       name,
					Hiden:      hiden,
				}
				catalogs = append(catalogs, catalog)
			}
		}
		return catalogs, nil
	}
}

// 搜尋目錄編號
func (catalogModel CatalogModel) Search(id int64) (Catalog []entities.Catalog, err error) {
	rows, err := catalogModel.Db.Query("select * from catalogs where catalog_id = ? and hiden=0", id)
	if err != nil {
		return nil, err
	} else {
		var catalogs []entities.Catalog
		for rows.Next() {
			var catalog_id int64
			var name string
			var hiden uint
			err2 := rows.Scan(&catalog_id, &name, &hiden)
			if err2 != nil {
				return nil, err2
			} else {
				catalog := entities.Catalog{
					Catalog_Id: catalog_id,
					Name:       name,
					Hiden:      hiden,
				}
				catalogs = append(catalogs, catalog)
			}
		}
		return catalogs, nil
	}
}

// 建立商品目錄
func (catalogModel CatalogModel) Create(Catalog *entities.Catalog) (err error) {
	result, err := catalogModel.Db.Exec("insert into catalogs(name, hiden) values(?, ?)",
		Catalog.Name, Catalog.Hiden)
	if err != nil {
		return err
	} else {
		Catalog.Catalog_Id, _ = result.LastInsertId()
		return nil
	}
}

// 更新商品目錄
func (catalogModel CatalogModel) Update(Catalog *entities.Catalog) (int64, error) {
	result, err := catalogModel.Db.Exec("update catalogs set name=?, hiden=? where catalog_id=?",
		Catalog.Name, Catalog.Hiden, Catalog.Catalog_Id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}

// 刪除商品目錄
func (catalogModel CatalogModel) Delete(id int64) (int64, error) {
	result, err := catalogModel.Db.Exec("delete from catalogs where catalog_id = ?", id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}
