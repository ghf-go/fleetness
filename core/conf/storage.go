package conf

type StorageConfig struct {
	Driver     string `yaml:"driver"`
	CdnHost    string `yaml:"cnd_host"`
	Ak         string `yaml:"ak"`
	Sk         string `yaml:"sk"`
	Bucket     string `yaml:"bucket"`
	UploadHost string `yaml:"upload_host"`
}
