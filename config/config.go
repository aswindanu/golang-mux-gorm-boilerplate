package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Charset  string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			// MYSQL
			// Dialect:  "mysql",
			// Host:     "db",
			// Port:     3306,
			// Username: "root",
			// Password: "password",
			// Name:     "todoapp",
			// Charset:  "utf8",

			// POSTGRES
			Dialect:  "postgres",
			Host:     "db",
			Port:     5432,
			Username: "root",
			Password: "password",
			Name:     "todoapp",
			Charset:  "utf8",
		},
	}
}
