package conf

type appConfig struct {
	Port  int  `yaml:"port"`
	Debug bool `yaml:"debug"`
}
