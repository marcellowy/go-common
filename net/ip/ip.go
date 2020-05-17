package ip

import (
	"math/rand"
	"net"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// 获取一个空闲的端口
// 如果没有就在范围内返回一个随机数 [57000-65530)
func GetIdlePort() (int, error) {

	var (
		b = 57000
		e = 65530
		c = e - b
	)

	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		//return 0, err
		return b + rand.Intn(c), nil
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return b + rand.Intn(c), nil
	}
	defer l.Close()

	return l.Addr().(*net.TCPAddr).Port, nil
}

// 获取本地IP列表
func LocalIPv4s() ([]string, error) {
	var ips []string
	address, err := net.InterfaceAddrs()
	if err != nil {
		return ips, err
	}

	for _, a := range address {
		if ipNet, ok := a.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			ips = append(ips, ipNet.IP.String())
		}
	}

	return ips, nil
}

// 获取一个本地的IP, 若本地没有IP或者获取失败返回127.0.0.1
func LocalIp() string {

	var def = "127.0.0.1"

	ips, err := LocalIPv4s()
	if err != nil {
		return def
	}
	return ips[0]
}

// 获取STKE实例IP, 仅支持腾讯的STKE实例
func LocalIpWithPod() string {
	var ip = os.Getenv("POD_IP")
	if ip == "" {
		panic("POD IP empty")
	}
	return ip
}
