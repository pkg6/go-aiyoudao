package aiyoudao

import "net/url"

type BodyMaps map[string][]string

func MergeBodyMaps(baseBodyMaps BodyMaps, bodyMaps ...BodyMaps) BodyMaps {
	for _, bodyMap := range bodyMaps {
		for s, strings := range bodyMap {
			baseBodyMaps[s] = strings
		}
	}
	return baseBodyMaps
}
func (m BodyMaps) Add(key, value string) {
	if _, ok := m[key]; !ok {
		m[key] = []string{value}
	}
}

func (m BodyMaps) UrlValues() url.Values {
	params := url.Values{}
	for k, v := range m {
		for pv := range v {
			params.Add(k, v[pv])
		}
	}
	return params
}
