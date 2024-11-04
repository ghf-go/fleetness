package conf

type dbConfig struct {
	DbName         string   `yaml:"dnmane"`
	MaxIdleCons    int      `yaml:"max_idle_cons"`
	MaxOpenCons    int      `yaml:"max_open_cons"`
	ConMaxIdleTime int      `yaml:"con_max_idle_time"`
	ConMaxLifeTime int      `yaml:"con_max_life_time"`
	Write          string   `yaml:"write"`
	Reads          []string `yaml:"reads"`
}
