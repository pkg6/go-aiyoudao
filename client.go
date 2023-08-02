package gaiyoudao

import (
	"encoding/base64"
	"github.com/google/uuid"
	"github.com/pkg6/go-requests"
	"io/ioutil"
	"net/url"
	"os"
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
		Debug     bool
	}
)

func NewSingleton(appKey, appSecret string) *Client {
	once.Do(func() {
		single = New(appKey, appSecret)
	})
	return single
}

func New(appKey, appSecret string) *Client {
	return &Client{AppKey: appKey, AppSecret: appSecret}
}

func (c *Client) SaveFile(path string, data []byte, needDecode bool) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if needDecode {
		data, _ = base64.StdEncoding.DecodeString(string(data))
	}
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) ReadFileAsBase64(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	fd, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	} else {
		return base64.StdEncoding.EncodeToString(fd), nil
	}
}

func (c *Client) BuildRequestBody(signType string, bodyMaps BodyMaps) url.Values {
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
	return params
}

func (c *Client) PostForm(path string, body url.Values, resp any) error {
	response, err := requests.Post(RootURI+path, body, func(client *requests.Client) {
		client.AsForm()
		if c.Debug {
			client.Debug()
		}
	})
	if err != nil {
		return err
	}
	_ = response.Unmarshal(&resp)
	return nil
}
