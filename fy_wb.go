package aiyoudao

//https://ai.youdao.com/DOCSIRMA/html/trans/api/wbfy/index.html

type WBfyResponseWF struct {
	Wf struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"wf"`
}

type WBfyResponseBasic struct {
	Phonetic   string           `json:"phonetic"`
	UkPhonetic string           `json:"uk-phonetic"`
	UsPhonetic string           `json:"us-phonetic"`
	UkSpeech   string           `json:"uk-speech"`
	UsSpeech   string           `json:"us-speech"`
	Explains   []string         `json:"explains"`
	ExamType   []string         `json:"exam_type"`
	Wfs        []WBfyResponseWF `json:"wfs"`
}
type WBfyResponseWeb struct {
	Key   string   `json:"key"`
	Value []string `json:"value"`
}
type WBfyResponseURI struct {
	Url string `json:"url"`
}
type WBfyResponse struct {
	ErrorCode       string            `json:"errorCode"`
	RequestId       string            `json:"requestId"`
	Query           string            `json:"query"`
	ReturnPhrase    []string          `json:"returnPhrase"`
	IsDomainSupport string            `json:"isDomainSupport"`
	Translation     []string          `json:"translation"`
	IsWord          bool              `json:"isWord"`
	MTerminalDict   WBfyResponseURI   `json:"mTerminalDict"`
	Basic           WBfyResponseBasic `json:"basic"`
	Web             []WBfyResponseWeb `json:"web"`
	Dict            WBfyResponseURI   `json:"dict"`
	Webdict         WBfyResponseURI   `json:"webdict"`
	L               string            `json:"l"`
	TSpeakUrl       string            `json:"tSpeakUrl"`
	SpeakUrl        string            `json:"speakUrl"`
}

func (c *Client) WBfy(q, from, to string, bodyMaps ...BodyMaps) (resp WBfyResponse, err error) {
	err = c.PostForm("v3", "/api", &resp, BodyMaps{
		"q":    {q},
		"from": {from},
		"to":   {to},
	}, bodyMaps...)
	return resp, err
}
