package conf

type Conf struct {
	App  *appConfig          `yaml:"app"`
	Dbs  map[string]dbConfig `yaml:"dbs"`
	Log  *LogConfig          `yaml:"LogConfig"`
	Meta MetaConf            `yaml:"meta"`
}
