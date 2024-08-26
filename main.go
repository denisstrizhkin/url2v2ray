package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
)

type VmessData struct {
	Address  string `json:"add"`
	AlterId  string `json:"aid"`
	Host     string `json:"host"`
	UserId   string `json:"id"`
	Net      string `json:"net"`
	Path     string `json:"path"`
	Port     int    `json:"port"`
	Info     string `json:"ps"`
	Security string `json:"scy"`
	TLS      string `json:"tls"`
	Type     string `json:"type"`
	TypeV    string `json:"v"`
}

func die(values ...interface{}) {
	fmt.Fprintln(os.Stderr, values...)
	os.Exit(1)
}

func parseVMess(data []byte) string {
	vmess := VmessData{}
	err := json.Unmarshal(data, &vmess)
	if err != nil {
		die("can't decode vmess data:", err)
	}
	fmt.Fprintf(os.Stderr, "vmess data:", vmess)
	return ""
}

func parseUrl(url *url.URL) string {
	scheme := url.Scheme
	data, err := base64.StdEncoding.DecodeString(url.Host)
	if err != nil {
		die("can't decode url:", err)
	}
	fmt.Fprintln(os.Stderr, "scheme:", scheme)
	switch scheme {
	case "vmess":
		return parseVMess(data)
	default:
		die("unknown scheme")
	}
	return ""
}

func main() {
	exe := os.Args[0]
	args := os.Args[1:]
	if len(args) != 1 {
		die(fmt.Sprintf("usage: %s <url>\n", exe))
	}
	url, err := url.Parse(args[0])
	if err != nil {
		die("malformed url:", err)
	}
	fmt.Fprintln(os.Stderr, "url:", url)
	config := parseUrl(url)
	fmt.Println(config)
}
