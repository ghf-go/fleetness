package conf

type cacheConfig struct {
	Host            string `yaml:"host"`
	UserName        string `yaml:"username"`
	Passwd          string `yaml:"passwd"`
	MinIdleConns    int    `yaml:"min_idle_cons"`
	MaxIdleConns    int    `yaml:"max_idle_cons"`
	MaxActiveConns  int    `yaml:"max_active_cons"`
	ConnMaxIdleTime int    `yaml:"con_max_idle_time"`
	ConnMaxLifetime int    `yaml:"con_max_life_time"`
}
