package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/thedevsaddam/gojsonq"
	"helloworld/constant"
	"io/ioutil"
	"net/http"
)

func CommonEq(funcName string, contractName string, contractAddress string, abi string, funcParam []interface{}) string {
	requestData := map[string]interface{}{
		"user":            constant.User,
		"contractName":    contractName,
		"contractAddress": contractAddress,
		"funcName":        funcName,
		"contractAbi":     json.RawMessage(abi),
		"funcParam":       funcParam,
		"groupId":         1,
		"useCns":          false,
		"useAes":          false,
		"cnsName":         contractName,
		"version":         "",
	}
	requestDataBytes, _ := json.Marshal(requestData)
	req, err := http.NewRequest(http.MethodPost, constant.User, bytes.NewBuffer(requestDataBytes))
	if err != nil {
		fmt.Println("创建HTTP请求错误:", err)
		return ""
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送HTTP请求错误:", err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应主体错误:", err)
		return ""
	}
	if err != nil {
		fmt.Println("读取响应主体错误:", err)
		return ""
	}
	return string(body)
}

func GetJsonVal(body string, key string) interface{} {
	data := gojsonq.New().JSONString(body)
	val := data.Find(key)
	if val == nil {
		return ""
	}
	return val
}
