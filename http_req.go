/*
http请求库
*/
package goutils

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

//通用响应结构
type Http_ResponseCommon struct {
	Errcode string `json:"errcode"`
}

//Post请求
//@addr:目标的url
//@datas:请求数据
func Http_PostJson(url string, datas interface{}) error {
	// data collection
	b, err := json.Marshal(datas)
	if err != nil {
		return errors.New("Invalid json format")
	}
	cli := http.Client{
		Timeout: 10 * time.Second, //default time out
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(b))
	if err != nil {
		return err
	}
	resp, err := cli.Do(req)
	if err != nil {
		return err
	}

	// analize response body
	respD, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("No response body")
	}
	resp.Body.Close()
	var respContent Http_ResponseCommon
	err = json.Unmarshal(respD, &respContent)
	if err != nil {
		return errors.New("Unrecognized formation to response body")
	}

	if respContent.Errcode == "00000" {
		return nil
	} else {
		return errors.New(respContent.Errcode)
	}
}

//Post请求 带GZIP压缩
//@addr:目标的url
//@datas:请求数据
func PostJsonGzip(url string, datas interface{}) error {
	// data collection
	b, err := json.Marshal(datas)
	if err != nil {
		return errors.New("Invalid json format")
	}

	// compress
	var gBuffer bytes.Buffer
	gWritter := gzip.NewWriter(&gBuffer)
	gWritter.Write(b)
	defer gWritter.Close()
	gWritter.Flush()
	if err != nil {
		return err
	}

	cli := http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(gBuffer.Bytes()))
	if err != nil {
		return err
	}
	resp, err := cli.Do(req)
	if err != nil {
		return err
	}

	//读取响应
	respD, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("No response body")
	}
	resp.Body.Close()

	var respContent Http_ResponseCommon
	err = json.Unmarshal(respD, &respContent)
	if err != nil {
		return errors.New("Unrecognized formation to response body")
	}

	if respContent.Errcode == "00000" {
		return nil
	} else {
		return errors.New(respContent.Errcode)
	}
}
