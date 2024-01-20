package aiyoudao

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg6/go-requests"
	"strconv"
	"sync"
	"time"
)

var (
	RootURI = "https://openapi.youdao.com"
	once    sync.Once
	single  *Client // 单例
)

type (
	Client struct {
		AppKey    string
		AppSecret string
		Request   *requests.Client
	}
)

func NewSingleton(appKey, appSecret string) *Client {
	once.Do(func() {
		single = New(appKey, appSecret)
	})
	return single
}

func New(appKey, appSecret string) *Client {
	return &Client{
		AppKey:    appKey,
		AppSecret: appSecret,
		Request:   requests.New(),
	}
}

func (c *Client) PostForm(signType, path string, resp any, defaultBody BodyMaps, otherMaps ...BodyMaps) error {
	bodyMaps := MergeBodyMaps(defaultBody, otherMaps...)
	curTime := strconv.FormatInt(time.Now().Unix(), 10)
	salt := uuid.New().String()
	var sign string
	switch signType {
	case "v4":
		sign = requests.Sha256(c.AppKey + salt + curTime + c.AppSecret)
	case "v3":
		qs := bodyMaps["q"]
		if qs == nil {
			qs = bodyMaps["img"]
		}
		var q string
		for i := range qs {
			q += qs[i]
		}
		inputFun := func(q string) string {
			str := []rune(q)
			strLen := len(str)
			if strLen <= 20 {
				return q
			} else {
				return string(str[:10]) + strconv.Itoa(strLen) + string(str[strLen-10:])
			}
		}
		sign = requests.Sha256(c.AppKey + inputFun(q) + salt + curTime + c.AppSecret)
	}
	params := bodyMaps.UrlValues()
	params.Add("appKey", c.AppKey)
	params.Add("salt", salt)
	params.Add("curtime", curTime)
	params.Add("signType", signType)
	params.Add("sign", sign)
	return c.Request.PostFormUnmarshal(context.Background(), RootURI+path, params, resp)
}
