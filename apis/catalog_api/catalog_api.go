package catalog_api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ek2061/goshop/config"
	"github.com/ek2061/goshop/entities"
	"github.com/ek2061/goshop/models"
	"github.com/gorilla/mux"
)

// 列出全部目錄的內容
func FindAll(response http.ResponseWriter, request *http.Request) {
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, "資料庫連接發生錯誤")
	} else {
		catalogModel := models.CatalogModel{
			Db: db,
		}
		catalogs, err2 := catalogModel.FindAll()
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, "商品目錄搜尋失敗")
		} else {
			respondWithJson(response, http.StatusOK, catalogs)
		}
	}
}

// 搜尋某個目錄id的內容
func Search(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	// 將字串轉成數字確認是否被植入sql代碼
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		respondWithError(response, http.StatusBadRequest, "id引數錯誤")
		return
	}

	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, "資料庫連接發生錯誤")
	} else {
		catalogModel := models.CatalogModel{
			Db: db,
		}
		catalogs, err2 := catalogModel.Search(id)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, "商品目錄搜尋失敗")
		} else if catalogs == nil {
			// 找不到目錄就回傳404
			respondWithError(response, http.StatusNotFound, "商品目錄不存在或被隱藏")
		} else {
			respondWithJson(response, http.StatusOK, catalogs)
		}
	}
}

// 創建商品目錄使用json
func Create(response http.ResponseWriter, request *http.Request) {
	var catalog entities.Catalog
	err := json.NewDecoder(request.Body).Decode(&catalog)

	// ⽬錄名稱必填
	if catalog.Name == "" {
		respondWithError(response, http.StatusBadRequest, "欄位不能為空")
		return
	}

	// 商品⽬錄預設不隱藏
	catalog.Hiden = 0

	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, "資料庫連接發生錯誤")
	} else {
		catalogModel := models.CatalogModel{
			Db: db,
		}
		err2 := catalogModel.Create(&catalog)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, "創建商品目錄失敗")
		} else {
			respondWithJson(response, http.StatusOK, "更新商品目錄完成")
		}
	}
}

// 更新商品目錄使用json
func Update(response http.ResponseWriter, request *http.Request) {
	var catalog entities.Catalog
	err := json.NewDecoder(request.Body).Decode(&catalog)

	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, "資料庫連接發生錯誤")
	} else {
		catalogModel := models.CatalogModel{
			Db: db,
		}
		_, err2 := catalogModel.Update(&catalog)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, "更新商品目錄失敗")
		} else {
			respondWithJson(response, http.StatusOK, catalog)
		}
	}
}

func Delete(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	// 將字串轉成數字確認是否被植入sql代碼
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		respondWithError(response, http.StatusBadRequest, "id引數錯誤")
		return
	}

	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, "資料庫連接發生錯誤")
	} else {
		catalogModel := models.CatalogModel{
			Db: db,
		}
		_, err2 := catalogModel.Delete(id)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, "刪除商品目錄失敗")
		} else {
			respondWithJson(response, http.StatusOK, "刪除商品目錄完成")
		}
	}
}

// 回覆錯誤碼
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

// 回覆成功結果使用JSON格式
func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
