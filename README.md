微定位技術居家自主健康管理手環 後端伺服器專案(Wristband NBIOT Server)
===========================
該文件用來說明如何使用手環後端伺服器的程式配置方式。


依照以下步驟去配置專案後，即可測試專案。
----
### 1. 部屬專案到伺服器
    下載完專案後將檔案放至 user/go/src/ 底下
### 2. 設定資料庫 
	創建資料庫 名稱為: wristband
	匯入資料庫: 透過專案裏頭附的檔案 wristband.sql 去匯入
### 3. 設定資料庫使用者 : wristband-nb-server/database/database.go
### 4. 測試專案: wristband-nb-server/ 路徑
	編譯程式: go build
	執行程式: wristband-nb-server.exe
	中斷程式: 按鍵盤上的 Ctrl + C
	檢查完沒問題後，打包程式成執行檔: go build
### 5. 最後使用 deamon 服務去執行專案
