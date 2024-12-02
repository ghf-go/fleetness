package conf

type Conf struct {
	App     *appConfig             `yaml:"app"`
	Dbs     map[string]dbConfig    `yaml:"dbs"`
	Cache   map[string]cacheConfig `yaml:"cache"`
	Log     *logConfig             `yaml:"LogConfig"`
	Stmp    map[string]smtpConfig  `yaml:"smtp"`
	Payment PaymentConfig          `yaml:"payment"`
	Lang    langConf               `yaml:"lang"`
	Storage storageConfig          `yaml:"storage"`
	Meta    MetaConf               `yaml:"meta"`
}
