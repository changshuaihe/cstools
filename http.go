package cstools

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

func HttpGet(url string, headers map[string]string) string {
	client := &http.Client{Timeout: 10 * time.Second}
	//生成要访问的url
	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)

	//增加header选项
	for k, v := range headers {
		reqest.Header.Add(k, v)
	}

	if err != nil {
		panic(err)
	}
	//处理返回结果
	response, _ := client.Do(reqest)
	defer response.Body.Close()
	respByte, _ := ioutil.ReadAll(response.Body)
	return string(respByte)
}

func HttpPost(url string, data string, headers map[string]string) string {
	client := &http.Client{Timeout: 10 * time.Second}

	reqest, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))
	//增加header选项
	for k, v := range headers {
		reqest.Header.Add(k, v)
	}
	if err != nil {
		panic(err)
	}
	defer reqest.Body.Close()

	resp, error := client.Do(reqest)
	if error != nil {
		panic(error)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}
