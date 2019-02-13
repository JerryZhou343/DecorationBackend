package config

type config struct {
	Port int `yaml:"port"`
	Log  logConfig
}

type logConfig struct {
	Size         int64  `yaml:"size"`
	DateFlag     bool   `yaml:"dateFlag"`
	Path         string `yaml:"path"`
	CompressFlag bool   `yaml:"compressFlag"`
	Name         string `yaml:"fileName"`
}
