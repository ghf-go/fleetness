package conf

type storageConfig struct {
	Driver    string `yaml:"driver"`
	CdnHost   string `yaml:"cnd_host"`
	Ak        string `yaml:"ak"`
	Sk        string `yaml:"sk"`
	BluckName string `yaml:"bluch_name"`
}
