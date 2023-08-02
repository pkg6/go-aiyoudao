package aiyoudao

import (
	"encoding/base64"
	"io"
)

var (
	DefaultDetectType = "10012"
)

//https://ai.youdao.com/DOCSIRMA/html/ocr/api/tyocr/index.html

type TYOcrResponseRegionsWords struct {
	BoundingBox string `json:"boundingBox"`
	Word        string `json:"word"`
}
type TYOcrResponseRegionsLines struct {
	BoundingBox string                      `json:"boundingBox"`
	Lang        string                      `json:"lang"`
	Style       string                      `json:"style"`
	Text        string                      `json:"text"`
	TextHeight  int                         `json:"text_height"`
	Words       []TYOcrResponseRegionsWords `json:"words"`
}

type TYOcrResponseRegions struct {
	BoundingBox string                      `json:"boundingBox"`
	Dir         string                      `json:"dir"`
	Lang        string                      `json:"lang"`
	Lines       []TYOcrResponseRegionsLines `json:"lines"`
}

type TYOcrResponse struct {
	ErrorCode string `json:"errorCode"`
	Result    struct {
		Orientation string                 `json:"orientation"`
		Regions     []TYOcrResponseRegions `json:"regions"`
	} `json:"Result"`
}

func (c *Client) TYOcrFile(file string, langType string, bodyMaps ...BodyMaps) (resp TYOcrResponse, err error) {
	q, err := ReadFileAsBase64(file)
	if err != nil {
		return resp, err
	}
	return c.TYOcr(q, langType, bodyMaps...)
}

func (c *Client) TYOcrForReader(reader io.Reader, langType string, bodyMaps ...BodyMaps) (resp TYOcrResponse, err error) {
	fd, err := io.ReadAll(reader)
	if err != nil {
		return resp, err
	}
	return c.TYOcr(base64.StdEncoding.EncodeToString(fd), langType, bodyMaps...)
}

func (c *Client) TYOcr(imgBase64, langType string, bodyMaps ...BodyMaps) (resp TYOcrResponse, err error) {
	bodyMap := MergeBodyMaps(BodyMaps{
		"img":       {imgBase64},
		"langType":  {langType},
		"docType":   {"json"},
		"imageType": {"1"},
	}, bodyMaps...)
	bodyMap.Add("detectType", DefaultDetectType)
	err = c.PostForm("v3", "/ocrapi", &resp, bodyMap)
	return resp, err
}
