package config

import (
	"errors"
	"github.com/RockkleyPushPost/common/database"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Database    *database.DBConfig `json:"database" yaml:"database"`
	JwtSecret   string             `json:"jwt_secret" yaml:"jwt_secret"`
	Server      ServerConfig       `json:"server" yaml:"server"`
	ServiceName string             `json:"serviceName" yaml:"serviceName"`
}

type ServerConfig struct {
	Host string `json:"host" yaml:"host"`
	Port string `json:"port" yaml:"port"`
}

func LoadYamlConfig(path string) (*Config, error) {
	if path == "" {

		return nil, errors.New("no yaml config path provided")
	}

	cfg := &Config{}

	file, err := os.ReadFile(path)

	if err != nil {

		return nil, err
	}

	replaced := os.ExpandEnv(string(file))
	err = yaml.Unmarshal([]byte(replaced), cfg)
	return cfg, err
	//decoder := yaml.NewDecoder(file)
	//
	//if err := decoder.Decode(&config); err != nil {
	//
	//	return nil, err
	//}
	//
	//if &config == nil {
	//
	//	return nil, errors.New("loaded services config is nil")
	//
	//}

	//return &config, nil
}

func (c *ServerConfig) Validate() error {
	if c.Host == "" {

		return errors.New("missing host")
	}

	if c.Port == "" {

		return errors.New("missing port")
	}

	return nil
}
