package main

import (
	"fmt"
	"github.com/pkg6/go-aiyoudao"
)

func main() {
	youdao := aiyoudao.NewSingleton("appKey", "appSecret")
	bfy, err := youdao.WBfy("你好", "zh-CHS", "en")
	fmt.Println(err)
	fmt.Println(bfy)
}
