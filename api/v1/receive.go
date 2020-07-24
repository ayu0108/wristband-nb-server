package v1

import (
	"encoding/json"
	"strconv"
	"wristband-nb-server/models"
)

// GetReceiveDatas -
func GetReceiveDatas() map[int]string {
	var receiveDatas = models.GetReceiveDatas()
	m := make(map[int]string)

	for id, receiveData := range receiveDatas {
		jsonStr, err := json.Marshal(receiveData)
		if err != nil {
			continue
		}
		m[id] = string(jsonStr)
	}

	return m
}

// AddReceiveData : 新增資料
func AddReceiveData(form []string) bool {
	SerialNumber, err := strconv.Atoi(form[3])
	if err != nil {
		return false
	}

	Singal, err := strconv.Atoi(form[4])
	if err != nil {
		return false
	}

	reciveData := models.ReceiveData{
		DeviceID:     form[0],
		BeaconName:   form[1],
		BeaconData:   form[2],
		SerialNumber: SerialNumber,
		Singal:       Singal,
		IPAddress:    form[5],
	}

	var result = models.AddReceiveData(reciveData)

	return result
}
