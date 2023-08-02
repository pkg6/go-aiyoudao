package aiyoudao

import (
	"encoding/base64"
	"io"
)

//https://ai.youdao.com/DOCSIRMA/html/learn/api/sjsxtcc/index.html

type OcrWritingEraseResponse struct {
	RequestId       string `json:"requestId"`
	ErrorCode       string `json:"errorCode"`
	OriginalImg     string `json:"originalImg"`
	EraseEnhanceImg string `json:"eraseEnhanceImg"`
}

func (c *Client) OcrWritingEraseForFile(file string, bodyMaps ...BodyMaps) (resp OcrWritingEraseResponse, err error) {
	q, err := ReadFileAsBase64(file)
	if err != nil {
		return resp, err
	}
	return c.OcrWritingErase(q, bodyMaps...)
}

func (c *Client) OcrWritingEraseForReader(reader io.Reader, bodyMaps ...BodyMaps) (resp OcrWritingEraseResponse, err error) {
	fd, err := io.ReadAll(reader)
	if err != nil {
		return resp, err
	}
	return c.OcrWritingErase(base64.StdEncoding.EncodeToString(fd), bodyMaps...)
}

// OcrWritingErase
//试卷手写体擦除
func (c *Client) OcrWritingErase(q string, bodyMaps ...BodyMaps) (resp OcrWritingEraseResponse, err error) {
	err = c.PostForm("v3", "/ocr_writing_erase", &resp, BodyMaps{
		"q": {q},
	}, bodyMaps...)
	return resp, err
}
