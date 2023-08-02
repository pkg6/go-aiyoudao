package gaiyoudao

//https://ai.youdao.com/DOCSIRMA/html/trans/api/plwbfy/index.html

type PLWBfyResponse struct {
	TranslateResults []struct {
		Query        string `json:"query"`
		Translation  string `json:"translation"`
		Type         string `json:"type"`
		VerifyResult string `json:"verifyResult"`
	} `json:"translateResults"`
	RequestId string `json:"requestId"`
	ErrorCode string `json:"errorCode"`
	L         string `json:"l"`
}

func (c *Client) PLWBfy(qs []string, from, to string, bodyMaps ...BodyMaps) (resp PLWBfyResponse, err error) {
	bodyMap := bodyMapsMerge(BodyMaps{
		"q":    qs,
		"from": {from},
		"to":   {to},
	}, bodyMaps...)
	if err = c.PostForm("/v2/api", c.BuildRequestBody("v3", bodyMap), &resp); err != nil {
		return resp, err
	}
	return resp, nil
}
