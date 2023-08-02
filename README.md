# 有道智云AI开放平台
## 安装

~~~
go get github.com/pkg6/gaiyoudao
~~~

## 基本使用

~~~
package main

import "github.com/pkg6/gaiyoudao"

func main() {
	youdao := gaiyoudao.NewSingleton("appKey", "appSecret")
	youdao.WBfy("你好", "zh-CHS", "en")
}

~~~

## 功能列表

| 功能                                                         | 方法                                                         |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| [批量文本翻译](https://ai.youdao.com/DOCSIRMA/html/trans/api/plwbfy/index.html) | PLWBfy                                                       |
| [图片翻译](https://ai.youdao.com/DOCSIRMA/html/trans/api/tpfy/index.html) | TPFy，TPFyForReader，TPFyForFile                             |
| [文本翻译](https://ai.youdao.com/DOCSIRMA/html/trans/api/wbfy/index.html) | WBfy                                                         |
| [试卷手写体擦除](https://ai.youdao.com/DOCSIRMA/html/learn/api/sjsxtcc/index.html) | OcrWritingErase，OcrWritingEraseForReader，OcrWritingEraseForFile |

