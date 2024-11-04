package conf

type Conf struct {
	App *appConfig          `yaml:"app"`
	Dbs map[string]dbConfig `yaml:"dbs"`
}
