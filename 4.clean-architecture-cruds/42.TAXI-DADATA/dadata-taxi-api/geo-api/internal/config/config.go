// internal/config/config.go
package config

type Config struct {
	DB struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
	}
	Server struct {
		Port string
	}
	JWT struct {
		Secret string
	}
	DaData struct {
		APIKey    string
		SecretKey string
	}
}

func Load() (*Config, error) {
	// Загрузка из env или config файла
}
