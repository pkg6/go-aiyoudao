# 有道智云AI开放平台
## 安装

~~~
go get github.com/pkg6/go-aiyoudao
~~~

## 基本使用

~~~
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
~~~

## 功能列表

| 功能                                                         | 方法                                                         |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| [批量文本翻译](https://ai.youdao.com/DOCSIRMA/html/trans/api/plwbfy/index.html) | PLWBfy                                                       |
| [图片翻译](https://ai.youdao.com/DOCSIRMA/html/trans/api/tpfy/index.html) | TPFy，TPFyForReader，TPFyForFile                             |
| [文本翻译](https://ai.youdao.com/DOCSIRMA/html/trans/api/wbfy/index.html) | WBfy                                                         |
| [试卷手写体擦除](https://ai.youdao.com/DOCSIRMA/html/learn/api/sjsxtcc/index.html) | OcrWritingErase，OcrWritingEraseForReader，OcrWritingEraseForFile |
| [通用OCR](https://ai.youdao.com/DOCSIRMA/html/ocr/api/tyocr/index.html) | TYOcr，TYOcrForReader，TYOcrFile                             |
| [语音合成](https://ai.youdao.com/DOCSIRMA/html/tts/api/yyhc/index.html) | YYHC                                                         |

