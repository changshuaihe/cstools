package lib

import "encoding/json"

type returnMsg struct {
	Message string      `json:"message"`
	Success string      `json:"success"`
	Data    interface{} `json:"data"`
}

func Success(msg string, data interface{}) string {
	res := returnMsg{}
	res.Message = msg
	res.Success = "true"
	res.Data = data
	returnByte, _ := json.Marshal(res)
	returnStr := string(returnByte)
	return returnStr
}

func Error(msg string) string {
	res := returnMsg{}
	res.Message = msg
	res.Data = nil
	res.Success = "false"
	returnByte, _ := json.Marshal(res)
	returnStr := string(returnByte)
	return returnStr
}
