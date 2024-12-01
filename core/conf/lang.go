package conf

import "strings"

type langConf map[string]map[string]string

// 生成多语言
func (c langConf) Lang(key, lang string) string {
	lang = strings.ToLower(lang)
	if r, ok := c[lang]; ok {
		if ret, isok := r[key]; isok {
			return ret
		}
		return ""
	}
	return c.Lang(key, "us-en")
}
