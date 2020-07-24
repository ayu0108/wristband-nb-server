package models

import (
	_ "github.com/jinzhu/gorm" // "github.com/jinzhu/gorm"
)

// ReceiveData - 接收資料型態
type ReceiveData struct {
	Model
	DeviceID     string `json:"device_id"`
	IPAddress    string `json:"ip_address"`
	BeaconName   string `json:"beacon_name"`
	BeaconData   string `json:"beacon_data"`
	SerialNumber int    `json:"serial_number"`
	Singal       int    `json:"singal"`
}

// TableName - 設定操作 Table 名稱
func (ReceiveData) TableName() string {
	return "receive_data"
}

// GetReceiveDatas - 取得資料
func GetReceiveDatas() []ReceiveData {
	var receiveDatas []ReceiveData

	db.Find(&receiveDatas)
	return receiveDatas
}

// AddReceiveData - 新增資料
func AddReceiveData(data ReceiveData) bool {
	err := db.Create(&data).Error

	if err != nil {
		return false
	}

	return true
}
