package util

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"iochen.com/v2gen/v2/common/base64"
)

func splitVmess(s string) []string {
	var sep = map[rune]bool{
		' ':  true,
		'\n': true,
		',':  true,
		';':  true,
		'\t': true,
		'\f': true,
		'\v': true,
		'\r': true,
	}
	return strings.FieldsFunc(s, func(r rune) bool {
		return sep[r]
	})
}

func GetVmessList(url string) ([]string, error) {
	resp, err := http.Get(url) //请求base64Vmess
	defer func(Body io.ReadCloser) { _ = Body.Close() }(resp.Body)
	if err != nil {
		return []string{}, err
	}
	base64Vmess, err := ioutil.ReadAll(resp.Body) //读取base64Vmess
	if err != nil {
		return nil, err
	}
	strVmess, err := base64.Decode(string(base64Vmess)) //解码base64Vmess为strVmess
	if err != nil {
		return nil, err
	}
	return splitVmess(strVmess), nil //分割strVmess
}
