package tools

import (
	"compress/gzip"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Client 客户端
type Client struct {
	path   string
	query  map[string]string
	header map[string]string
	host   string
	port   uint

	request  *http.Request
	response *http.Response
	// 证书
	certPool    *x509.CertPool
	certificate *tls.Certificate

	// 请求超时时间
	duration time.Duration
}

// NewClient 创建新 Client
func NewClient() *Client {
	return &Client{}
}

// SetTimeout 设置超时时间
func (c *Client) SetTimeout(duration time.Duration) {
	c.duration = duration
}

// SetHeader 设置header
func (c *Client) SetHeader(key string, value string) {
	if c.header == nil {
		c.header = make(map[string]string)
	}

	if len(key) > 0 && len(value) > 0 {
		c.header[key] = value
	}
}

// SetCert 设置证书
func (c *Client) SetCert(crt, key, ca []byte) (err error) {

	c.certPool = x509.NewCertPool()
	c.certPool.AppendCertsFromPEM(ca)
	if *c.certificate, err = tls.X509KeyPair(crt, key); err != nil {
		return
	}
	return
}

// Post 发送POST请求
func (c *Client) Post(url string, requestBody string) ([]byte, error) {
	var (
		err error
	)
	if c.request, err = http.NewRequest("POST", url, strings.NewReader(requestBody)); err != nil {
		return nil, err
	}
	return c.do()
}

// Get 发送GET请求
func (c *Client) Get(url string) ([]byte, error) {
	var (
		err error
	)
	if c.request, err = http.NewRequest("GET", url, nil); err != nil {
		return nil, err
	}
	return c.do()
}

// do 发送请求
func (c *Client) do() ([]byte, error) {

	var (
		reader io.ReadCloser
		body   []byte
		err    error
	)

	// 设置压缩
	if c.header == nil {
		c.header = make(map[string]string)
	}

	// 设置可以接受gzip压缩
	c.header["Accept-Encoding"] = "gzip, deflate"

	// 设置头部
	for k, v := range c.header {
		c.request.Header.Set(k, v)
	}

	client := http.Client{}

	if c.certificate != nil {
		// 双向证书验证
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				ClientCAs:          c.certPool,
				Certificates:       []tls.Certificate{*c.certificate},
				InsecureSkipVerify: false, // 不验证服务端, 根据经验以前验证服务端会偶尔出错
			},
		}
	}

	// 设置超时时间
	if c.duration == 0 {
		// 默认超时 30s
		c.duration = time.Second * 10
	}
	client.Timeout = c.duration

	// 发送请求
	if c.response, err = client.Do(c.request); err != nil {
		return nil, err
	}
	defer c.response.Body.Close()

	if c.response.StatusCode != 200 {
		return nil, fmt.Errorf("client do failed; response status code: %d", c.response.StatusCode)
	}

	// 检查是否被服务器压缩, 被压缩需要解压
	if c.response.Header.Get("Content-Encoding") == "gzip" {
		// 解压缩
		if reader, err = gzip.NewReader(c.response.Body); err != nil {
			return nil, err
		}
	} else {
		reader = c.response.Body
	}

	if body, err = ioutil.ReadAll(reader); err != nil {
		return nil, err
	}

	return body, nil
}
