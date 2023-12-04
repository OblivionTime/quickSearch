package config

type System struct {
	Addr         int    `mapstructure:"addr" json:"addr" yaml:"addr"`                            // 端口值
	RouterPrefix string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"` //  路由全局前缀
}
