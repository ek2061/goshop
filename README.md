# goshop
簡單的商品資料和目錄api

## 如何開啟
```bash
$ cd goshop
$ go mod download
$ go run main.go
```

## 回傳引數說明
* **products 商品資料**  
"Product_Id": 商品id，主鍵  
"Catalog_Id": 目錄id，外鍵  
"Name": 商品名稱  
"Cost": 成本  
"Price": 價格  
"Description": 描述  
"On_Sale": 是否上架，0為false，1為true，預設0(不上架)  
"Start_Sell": 開始販售時間  
"End_Sell": 結束販售時間  
*"Catalog_Name": 目錄名稱，會透過關聯得到*  

* **catalogs 商品目錄**  
"Catalog_Id": 目錄id，主鍵  
"Name": 目錄名稱  
"Hiden": 是否隱藏，0為false，1為true，預設0(不隱藏)  

* **發生錯誤**  
成功: 回傳200  
失敗: 回傳400  
找不到商品資料: 回傳404  

## api說明
### 商品資料api  
* **GET /api/product**  
會返回所有商品資料，只要條件符合
1. 現在時間在販售期間內
2. 正在上架  

* **GET /api/product/:id**  
會返回該id的商品資料，只要條件符合
1. 現在時間在販售期間內
2. 正在上架  
3. id引數是合法數字  

* **POST /api/product**  
送出json，格式如回傳引數說明，只要條件符合  
1. 商品名稱、商品說明、商品開始販售時間、商品結束販售時間必填  
2. 商品id和目錄名稱不用填  
3. 商品售價 > 商品成本  
4. 商品開始販售時間 <= 商品結束販售時間  

* **PUT /api/product**  
送出json，格式如回傳引數說明，只要條件符合  
1. 商品id、商品名稱、商品說明、商品開始販售時間、商品結束販售時間必填  
2. 目錄名稱不用填  
3. 商品售價 > 商品成本  
4. 商品開始販售時間 <= 商品結束販售時間  

* **DELETE /api/product/:id**  
會刪除該id的商品資料，只要條件符合  
1. id引數是合法數字  
2. 刪除的id存在  

### 商品目錄api  
* **GET /api/catalog**  
會返回所有商品目錄，只要條件符合
1. 目前不隱藏  

* **GET /api/catalog/:id**  
會返回該id的商品目錄，只要條件符合
1. 目前不隱藏  
2. id引數是合法數字     

* **POST /api/catalog**  
送出json，格式如回傳引數說明，只要條件符合  
1. ⽬錄名稱、是否隱藏必填  
2. 目錄id不用填  

* **PUT /api/catalog**  
送出json，格式如回傳引數說明，只要條件符合  
1. 目錄id、⽬錄名稱、是否隱藏必填  

* **DELETE /api/catalog/:id**  
會刪除該id的商品目錄，只要條件符合  
1. id引數是合法數字  
2. 刪除的id存在  
