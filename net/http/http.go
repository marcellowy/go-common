package http

// Method
type Method string

// Version
type Version string

const (
	MethodPost Method = "POST"
	MethodGet  Method = "GET"

	Version10 Version = "HTTP/1.0"
	Version11 Version = "HTTP/1.1"
	Version20 Version = "HTTP1.1"
)

// Client
type Client struct {
	method  Method
	path    string
	query   map[string]string
	version Version
	header  map[string]string
	host    string
	port    uint
}

//type Client struct {
//	header map[string]string
//
//	request  *netHttp.Request
//	response *netHttp.Response
//
//	// 证书
//	certPool    *x509.CertPool
//	certificate *tls.Certificate
//
//	// 请求超时
//	duration time.Duration
//}
//
//// 设置超时时间
//func (c *Client) SetTimeout(duration time.Duration) {
//	c.duration = duration
//}
//
//// 设置header
//func (c *Client) SetHeader(key string, value string) {
//
//	if c.header == nil {
//		c.header = make(map[string]string)
//	}
//
//	if len(key) > 0 && len(value) > 0 {
//		c.header[key] = value
//	}
//}
//
//// 设置证书
//func (c *Client) SetCert(crt, key, ca []byte) (err error) {
//
//	c.certPool = x509.NewCertPool()
//	c.certPool.AppendCertsFromPEM(ca)
//
//	if *c.certificate, err = tls.X509KeyPair(crt, key); err != nil {
//		return
//	}
//	return
//}
//
//// 发送POST请求
//func (c *Client) Post(url string, requestBody string) (responseBody string, err error) {
//	if c.request, err = netHttp.NewRequest("POST", url, strings.NewReader(requestBody)); err != nil {
//		return
//	}
//	return c.do()
//}
//
//// 发送GET请求
//func (c *Client) Get(url string) (responseBody string, err error) {
//	if c.request, err = netHttp.NewRequest("GET", url, strings.NewReader("")); err != nil {
//		return
//	}
//	return c.do()
//}
//
//// Send
//func (c *Client) do() (responseBody string, err error) {
//
//	// 设置压缩
//	if c.header == nil {
//		c.header = make(map[string]string)
//	}
//
//	// 设置可以接受gzip压缩
//	c.header["Accept-Encoding"] = "gzip, deflate"
//
//	// 设置头部
//	for k, v := range c.header {
//		c.request.Header.Set(k, v)
//	}
//
//	client := netHttp.Client{}
//
//	if c.certificate != nil {
//		// 双向证书验证
//		client.Transport = &netHttp.Transport{
//			TLSClientConfig: &tls.Config{
//				ClientCAs:          c.certPool,
//				Certificates:       []tls.Certificate{*c.certificate},
//				InsecureSkipVerify: false, // 不验证服务端, 根据经验以前验证服务端会偶尔出错
//			},
//		}
//	}
//
//	// 设置超时时间
//	if c.duration == 0 {
//		// 默认超时 30s
//		c.duration = time.Second * 30
//	}
//	client.Timeout = c.duration
//
//	// 发送请求
//	if c.response, err = client.Do(c.request.WithContext(context.TODO())); err != nil {
//		return
//	}
//	defer c.response.Body.Close()
//
//	if c.response.StatusCode != 200 {
//		// 与服务器通信失败
//		err = errors.New(fmt.Sprintf("client do failed; response status code: %d", c.response.StatusCode))
//		return
//	}
//
//	var (
//		reader io.ReadCloser
//		body   []byte
//	)
//
//	// 检查是否被服务器压缩, 被压缩需要解压
//	if c.response.Header.Get("Content-Encoding") == "gzip" {
//		// 解压缩
//		if reader, err = gzip.NewReader(c.response.Body); err != nil {
//			return
//		}
//	} else {
//		reader = c.response.Body
//	}
//
//	if body, err = ioutil.ReadAll(reader); err != nil {
//		return
//	}
//	responseBody = string(body)
//
//	return
//}
