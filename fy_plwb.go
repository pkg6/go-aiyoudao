package aiyoudao

//https://ai.youdao.com/DOCSIRMA/html/trans/api/plwbfy/index.html

type PLWBfyResponseTranslateResults struct {
	Query        string `json:"query"`
	Translation  string `json:"translation"`
	Type         string `json:"type"`
	VerifyResult string `json:"verifyResult"`
}

type PLWBfyResponse struct {
	TranslateResults []PLWBfyResponseTranslateResults `json:"translateResults"`
	RequestId        string                           `json:"requestId"`
	ErrorCode        string                           `json:"errorCode"`
	L                string                           `json:"l"`
}

func (c *Client) PLWBfy(qs []string, from, to string, bodyMaps ...BodyMaps) (resp PLWBfyResponse, err error) {
	err = c.PostForm("v3", "/v2/api", &resp, BodyMaps{
		"q":    qs,
		"from": {from},
		"to":   {to},
	}, bodyMaps...)
	return resp, err
}
