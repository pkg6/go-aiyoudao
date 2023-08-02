package main

import "github.com/pkg6/gaiyoudao"

func main() {
	youdao := gaiyoudao.NewSingleton("appKey", "appSecret")
	youdao.WBfy("你好", "zh-CHS", "en")
}
