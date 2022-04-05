package cstools

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func HttpGet(url string, headers map[string]string) (string, error) {
	client := &http.Client{Timeout: 300 * time.Second}
	//生成要访问的url
	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)

	//增加header选项
	for k, v := range headers {
		reqest.Header.Add(k, v)
	}

	if err != nil {
		fmt.Println(url)
		panic(err)
		return "", err
	}
	//处理返回结果
	response, err := client.Do(reqest)
	if err != nil {
		return "", err
	}
	respByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(respByte), err
}

func HttpPost(url string, data string, headers map[string]string) (string, error) {
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
		return "", err
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result), err
}
