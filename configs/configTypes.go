package configs

// Add config types here you need...

type AppConfig struct {
	Name        string `yaml:"name" toml:"name" json:"name"`
	Version     string `yaml:"version" toml:"version" json:"version"`
	Description string `yaml:"description" toml:"description" json:"description"`
	Environment string `yaml:"environment" toml:"environment" json:"environment"`
	LogLevel    string `yaml:"log-level" toml:"log-level" json:"log-level"`
}

type ApiServerConfig struct {
	BasePath   string `yaml:"base-path" toml:"base-path" json:"basePath"`
	ServerPort string `yaml:"server-port" toml:"base-path" json:"serverPort"`
	Cors       bool   `yaml:"cors" toml:"cors" json:"cors"`
}

type DbConfig struct {
	Driver string `yaml:"driver" toml:"driver" json:"driver"`
	Usr    string `yaml:"user" toml:"user" json:"user"`
	Pwd    string `yaml:"password" toml:"password" json:"password"`
	Host   string `yaml:"host" toml:"host" json:"host"`
	Port   string `yaml:"port" toml:"port" json:"port"`
	DbName string `yaml:"database-name" toml:"database-name" json:"databaseName"`
	Client string `yaml:"client" toml:"client" json:"client"`
}
