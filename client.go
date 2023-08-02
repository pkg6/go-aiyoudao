package aiyoudao

import (
	"context"
	"github.com/google/uuid"
	"github.com/zzqqw/gclient"
	"net/url"
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
		AppKey     string
		AppSecret  string
		HttpClient *gclient.Client
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
		AppKey:     appKey,
		AppSecret:  appSecret,
		HttpClient: gclient.New(),
	}
}

func (c *Client) PostForm(signType, path string, resp any, defaultBody BodyMaps, otherMaps ...BodyMaps) error {
	bodyMaps := MergeBodyMaps(defaultBody, otherMaps...)
	curTime := strconv.FormatInt(time.Now().Unix(), 10)
	salt := uuid.New().String()
	var sign string
	switch signType {
	case "v4":
		sign = gclient.Sha256String(c.AppKey + salt + curTime + c.AppSecret)
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
		sign = gclient.Sha256String(c.AppKey + inputFun(q) + salt + curTime + c.AppSecret)
	}
	params := url.Values{}
	for k, v := range bodyMaps {
		for pv := range v {
			params.Add(k, v[pv])
		}
	}
	params.Add("appKey", c.AppKey)
	params.Add("salt", salt)
	params.Add("curtime", curTime)
	params.Add("signType", signType)
	params.Add("sign", sign)
	return c.HttpClient.PostFormUnmarshal(context.Background(), RootURI+path, params, resp)
}
