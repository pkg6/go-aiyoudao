package gaiyoudao

//https://ai.youdao.com/DOCSIRMA/html/trans/api/wbfy/index.html

type WBfyResponse struct {
	ErrorCode       string   `json:"errorCode"`
	Query           string   `json:"query"`
	IsDomainSupport string   `json:"isDomainSupport"`
	Translation     []string `json:"translation"`
	Basic           struct {
		Phonetic   string   `json:"phonetic"`
		UkPhonetic string   `json:"uk-phonetic"`
		UsPhonetic string   `json:"us-phonetic"`
		UkSpeech   string   `json:"uk-speech"`
		UsSpeech   string   `json:"us-speech"`
		Explains   []string `json:"explains"`
	} `json:"basic"`
	Web []struct {
		Key   string   `json:"key"`
		Value []string `json:"value"`
	} `json:"web"`
	Dict struct {
		Url string `json:"url"`
	} `json:"dict"`
	Webdict struct {
		Url string `json:"url"`
	} `json:"webdict"`
	L         string `json:"l"`
	TSpeakUrl string `json:"tSpeakUrl"`
	SpeakUrl  string `json:"speakUrl"`
}

func (c *Client) WBfy(q, from, to string, bodyMaps ...BodyMaps) (resp WBfyResponse, err error) {
	bodyMap := bodyMapsMerge(BodyMaps{
		"q":    {q},
		"from": {from},
		"to":   {to},
	}, bodyMaps...)
	if err = c.PostForm("/api", c.BuildRequestBody("v3", bodyMap), &resp); err != nil {
		return resp, err
	}
	return resp, nil
}
