package config

type Config struct {
	ENV string `yaml:"env"`
	StoragePath string `yaml:"storage_path"`
	HTTPServer struct {
		Address string `yaml:"address"`
	}
}