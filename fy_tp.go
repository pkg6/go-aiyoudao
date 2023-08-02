package aiyoudao

import (
	"encoding/base64"
	"io"
)

//https://ai.youdao.com/DOCSIRMA/html/trans/api/tpfy/index.html

type TPFyResponseResRegions struct {
	BoundingBox string `json:"boundingBox"`
	LinesCount  int    `json:"linesCount"`
	Lineheight  int    `json:"lineheight"`
	Context     string `json:"context"`
	Linespace   int    `json:"linespace"`
	TranContent string `json:"tranContent"`
}

type TPFyResponse struct {
	Orientation string                   `json:"orientation"`
	LanFrom     string                   `json:"lanFrom"`
	TextAngle   string                   `json:"textAngle"`
	ErrorCode   string                   `json:"errorCode"`
	LanTo       string                   `json:"lanTo"`
	ResRegions  []TPFyResponseResRegions `json:"resRegions"`
}

func (c *Client) TPFyForFile(file, from, to string, bodyMaps ...BodyMaps) (resp TPFyResponse, err error) {
	q, err := ReadFileAsBase64(file)
	if err != nil {
		return resp, err
	}
	return c.TPFy(q, from, to, bodyMaps...)
}

func (c *Client) TPFyForReader(reader io.Reader, from, to string, bodyMaps ...BodyMaps) (resp TPFyResponse, err error) {
	fd, err := io.ReadAll(reader)
	if err != nil {
		return resp, err
	}
	return c.TPFy(base64.StdEncoding.EncodeToString(fd), from, to, bodyMaps...)
}

func (c *Client) TPFy(q, from, to string, bodyMaps ...BodyMaps) (resp TPFyResponse, err error) {
	err = c.PostForm("v3", "/ocrtransapi", &resp, BodyMaps{
		"q":    {q},
		"from": {from},
		"to":   {to},
	}, bodyMaps...)
	return resp, err
}
