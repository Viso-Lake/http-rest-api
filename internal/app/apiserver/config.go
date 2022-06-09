package apiserver

// Config ...
type Config struct {
	// BindAddr адрес на котором будет запущен сервер
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8000",
		LogLevel: "debug",
	}
}