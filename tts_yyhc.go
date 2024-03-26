package aiyoudao

//https://ai.youdao.com/DOCSIRMA/html/tts/api/yyhc/index.html

func (c *Client) YYHC(q, voiceName string, bodyMaps ...BodyMaps) ([]byte, error) {
	bodyMap := MergeBodyMaps(BodyMaps{
		"q":         {q},
		"voiceName": {voiceName},
		"format":    {"mp3"},
	}, bodyMaps...)
	return c.PostFormBinary("v3", "/ttsapi", bodyMap)
}
