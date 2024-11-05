package conf

type MetaConf map[string]any

func (m MetaConf) Get(key string) any {
	if r, ok := m[key]; ok {
		return r
	}
	return nil
}
