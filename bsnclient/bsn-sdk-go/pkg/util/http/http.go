package http

import (
	"bsn-sdk-go/pkg/common/errors"
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"github.com/wonderivan/logger"
	"io/ioutil"
	"net/http"
	"strings"
)

func SendPost(dataBytes []byte, url string, cert string) ([]byte, error) {

	var client *http.Client

	isHttps :=strings.Contains(url,"https://")

	if isHttps{
		logger.Debug("cert:",cert)
		if cert == "" {
			return nil,errors.New("HTTPS certificate not set")
		}

		//dirPath, err := filepath.Abs(".")
		//if err != nil {
		//	logger.Error("get current directory failed：", err.Error())
		//	return nil, err
		//}
		//读取https证书内容
		caCert, err := ioutil.ReadFile(cert)
		if err != nil {
			logger.Error("read HTTPS certificate content failed：", err.Error())
			return nil, err
		}
		//构建证书池
		caCertPool := x509.NewCertPool()
		//将读取的https证书内容添加到证书池
		caCertPool.AppendCertsFromPEM(caCert)
		//构建http请求客户端
		client = &http.Client{
			//定义单个HTTP请求的机制
			Transport: &http.Transport{
				//定义TLS客户端配置
				TLSClientConfig: &tls.Config{
					//添加RootCA证书池（此处将https的公钥证书添加到RootCA证书池中）
					RootCAs: caCertPool,
				},
			},
		}
	}else {
		logger.Debug("Http")
		tr := new(http.Transport)
		client = &http.Client{
			//定义单个HTTP请求的机制
			Transport: tr,
		}
	}
	//调用接口
	logger.Debug("request message：", string(dataBytes))
	response, err := client.Post(url, "application/json", bytes.NewReader(dataBytes))
	if err != nil {
		logger.Error("request failed：", err.Error())
		return nil, err
	}
	//从响应对象获取响应报文数据，并进行读取
	allBytes := []byte{}
	//缓冲区
	bytes := make([]byte, response.ContentLength)
	i, err := response.Body.Read(bytes)
	allBytes = append(allBytes, bytes[:i]...)

	for {
		i, err = response.Body.Read(bytes)
		if i == 0 {
			break
		}
		allBytes = append(allBytes, bytes[:i]...)
	}
	response.Body.Close()
	logger.Debug("response message：", string(allBytes))
	return allBytes, nil
}
