package main

import (
	"encoding/base64"
	"fmt"

	"github.com/skip2/go-qrcode"
	// "encoding/base64"
)

func main() { // 用来生成base64的二维码
	fmt.Println("test")
	var png []byte
	png, err := qrcode.Encode("http://baidu.com/", qrcode.Medium, 256)
	if err != nil {
		fmt.Println("error")
	}
	t := base64.StdEncoding.EncodeToString([]byte(png))
	fmt.Print(t)
}
