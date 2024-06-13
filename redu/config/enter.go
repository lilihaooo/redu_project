package config

// Config 配置文件
type Config struct {
	Mysql  `yaml:"mysql"`
	Redis  `yaml:"redis"`
	Logger `yaml:"logger"`
	Server `yaml:"server"`
	Es     `yaml:"es"`
}

// ResMap 错误映射
type ResMap map[int]string
