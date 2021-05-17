package product_api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/ek2061/goshop/config"
	"github.com/ek2061/goshop/entities"
	"github.com/ek2061/goshop/models"
	"github.com/gorilla/mux"
)

// 列出全部商品的內容
func FindAll(response http.ResponseWriter, request *http.Request) {
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, "資料庫連接發生錯誤")
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		products, err2 := productModel.FindAll()
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, "商品資料搜尋失敗")
		} else {
			respondWithJson(response, http.StatusOK, products)
		}
	}
}

// 搜尋某個商品id的內容
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
		productModel := models.ProductModel{
			Db: db,
		}
		products, err2 := productModel.Search(id)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, "商品資料搜尋失敗")
		} else if products == nil {
			// 找不到商品資料就回傳404
			respondWithError(response, http.StatusNotFound, "商品資料不存在或未上架")
		} else {
			respondWithJson(response, http.StatusOK, products)
		}
	}
}

// 創建商品使用json
func Create(response http.ResponseWriter, request *http.Request) {
	var product entities.Product
	err := json.NewDecoder(request.Body).Decode(&product)

	// ⽬錄名稱、商品名稱、商品說明、商品開始販售時間、商品結束販售時間必填
	if product.Catalog_Id == 0 ||
		product.Name == "" ||
		product.Description == "" ||
		product.Start_Sell == "" ||
		product.End_Sell == "" {
		respondWithError(response, http.StatusBadRequest, "欄位不能為空")
		return
	}

	// 商品售價 > 商品成本
	if product.Price <= product.Cost || product.Price == 0 || product.Cost == 0 {
		respondWithError(response, http.StatusBadRequest, "商品售價與成本標示錯誤")
		return
	}

	// 商品開始販售時間 <= 商品結束販售時間
	t1, err := time.Parse("2006-01-02 15:04:05", product.Start_Sell)
	t2, err := time.Parse("2006-01-02 15:04:05", product.End_Sell)
	if err == nil && t2.Before(t1) {
		respondWithError(response, http.StatusBadRequest, "商品販售時間設定有誤")
		return
	}

	// 商品預設不上架
	// product.On_Sale = 0

	// 接form值
	// catalog_id = request.FormValue("catalog_id")

	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, "資料庫連接發生錯誤")
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		err2 := productModel.Create(&product)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, "創建商品資料失敗")
		} else {
			respondWithJson(response, http.StatusOK, product)
		}
	}
}

// 更新商品資料使用json
func Update(response http.ResponseWriter, request *http.Request) {
	var product entities.Product
	err := json.NewDecoder(request.Body).Decode(&product)

	// 商品售價 > 商品成本
	if product.Price <= product.Cost || product.Price == 0 || product.Cost == 0 {
		respondWithError(response, http.StatusBadRequest, "商品售價與成本標示錯誤")
		return
	}

	// 商品開始販售時間 <= 商品結束販售時間
	t1, err := time.Parse("2006-01-02 15:04:05", product.Start_Sell)
	t2, err := time.Parse("2006-01-02 15:04:05", product.End_Sell)
	if err == nil && t2.Before(t1) {
		respondWithError(response, http.StatusBadRequest, "商品販售時間設定有誤")
		return
	}

	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, "資料庫連接發生錯誤")
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		n, err2 := productModel.Update(&product)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else if n == 0 {
			respondWithError(response, http.StatusBadRequest, "無法更新資料，可能是不存在或是不需要更新")
		} else {
			respondWithJson(response, http.StatusOK, "更新商品資料完成")
		}
	}
}

// 刪除商品資料
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
		productModel := models.ProductModel{
			Db: db,
		}
		n, err2 := productModel.Delete(id)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else if n == 0 {
			respondWithError(response, http.StatusBadRequest, "刪除的商品資料不存在")
		} else {
			respondWithJson(response, http.StatusOK, "刪除商品資料完成")
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
