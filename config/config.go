package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Api            Api
	Client         Client
	DB             DBConnection
	Logs           Logs
}

type Api struct {
	Port int
}

type Client struct {
	MaxSessionRetries           int `yaml:"max_session_retries"`
	SessionRetryIntervalSeconds int `yaml:"session_retry_interval_seconds"`
}

type DBConnection struct {
	Name           string `yaml:"name"`
	User           string `yaml:"user"`
	Password       string `yaml:"password"`
	Host           string `yaml:"host"`
	Port           int    `yaml:"port"`
	Type           string `yaml:"type"`
	TimeoutSeconds int    `yaml:"timeout_seconds"`
	SSLMode        string `yaml:"ssl_mode"`
	EncryptionKey  string `yaml:"encryption_key"`
}

type Logs struct {
	Access    string
	App       string
	AppLevel  string  `yaml:"app_level"`
}

// ReadFromConfig takes the path of a YAML config file and returns a Config struct
func ReadFromConfig(path *string) (*Config, error) {
	file, err := os.Open(*path)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	c := Config{}

	err = yaml.Unmarshal([]byte(data), &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (d DBConnection) BuildConnectionURL() string {
	return fmt.Sprintf("%s://%s:%s@%s:%d/%s?connect_timeout=%d&sslmode=%s", d.Type, d.User, d.Password, d.Host, d.Port, d.Name, d.TimeoutSeconds, d.SSLMode)
}
