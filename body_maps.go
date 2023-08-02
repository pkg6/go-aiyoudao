package aiyoudao

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
