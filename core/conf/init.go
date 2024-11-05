package conf

type Conf struct {
	App     *appConfig             `yaml:"app"`
	Dbs     map[string]dbConfig    `yaml:"dbs"`
	Cache   map[string]cacheConfig `yaml:"cache"`
	Log     *LogConfig             `yaml:"LogConfig"`
	Stmp    map[string]smtpConfig  `yaml:"smtp"`
	Payment PaymentConfig          `yaml:"payment"`
	Meta    MetaConf               `yaml:"meta"`
}
