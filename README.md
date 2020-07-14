微定位技術居家自主健康管理手環 後端伺服器專案(Wristband NBIOT Server)
===========================
該文件用來說明如何使用手環後端伺服器的程式配置方式。


依照以下步驟去配置專案後，即可測試專案。
----
### 1. 部屬專案到伺服器
    下載完專案後將檔案放至 user/go/src/ 底下
### 2. 設定資料庫
	創建資料庫 名稱為: nb_schema
	匯入資料庫: 透過專案裏頭附的檔案 nb_schema.sql 去匯入
### 3. 調整程式 mysql/mysql.go
	確認 database 變數設定像是: userName、password
	調整完程式後請先切換到該程式資料夾的路徑下，然後下指令 go install
### 4. 調整程式 tcpserver/tcpserver.go 
	檢查 port
	調整完一樣切換到程式資料夾的路徑下，然後下指令 go intsall
### 5. 測試專案: 
	切換路徑: cd /webservice
	執行程式: go run webservice.go
	中斷程式: 按鍵盤上的 `Ctrl + C`
	檢查完沒問題後，打包程式成執行檔: go build
### 6. 最後使用 deamon 服務去執行專案。
