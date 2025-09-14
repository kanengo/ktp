package transport

import (
	"fmt"
	"net"
	"strings"
)

func GetLocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok {
			ip := ipnet.IP.To4()
			if ip == nil {
				continue // 不是 IPv4
			}
			if !isValidBindableIPv4(ip) {
				continue
			}
			return ip.String(), nil
		}
	}

	return "", fmt.Errorf("未找到合适的本地 IPv4 地址")
}

func isValidBindableIPv4(ip net.IP) bool {
	// 回环 127.0.0.1
	if ip.IsLoopback() {
		return false
	}

	// 未指定 0.0.0.0
	if ip.IsUnspecified() {
		return false
	}

	// 多播地址 224.0.0.0 ~ 239.255.255.255
	if ip[0] >= 224 && ip[0] <= 239 {
		return false
	}

	// Link-local 169.254.x.x
	if ip[0] == 169 && ip[1] == 254 {
		return false
	}

	// 保留地址段 240.0.0.0+
	if ip[0] >= 240 {
		return false
	}

	// 示例地址
	if (ip[0] == 192 && ip[1] == 0 && ip[2] == 2) ||
		(ip[0] == 198 && ip[1] == 51 && ip[2] == 100) ||
		(ip[0] == 203 && ip[1] == 0 && ip[2] == 113) {
		return false
	}

	return true
}

func ContentSubtype(contentType string) string {
	left := strings.Index(contentType, "/")
	if left == -1 {
		return ""
	}
	right := strings.Index(contentType, ";")
	if right == -1 {
		right = len(contentType)
	}
	if right < left {
		return ""
	}
	return contentType[left+1 : right]
}
