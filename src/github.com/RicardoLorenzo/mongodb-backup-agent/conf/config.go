package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ConfigError struct {
	message string
	err     error
}

func (e *ConfigError) Error() string {
	return e.message
}

type Config struct {
	path       string
	properties map[string]string
}

func NewConfig() (Config, error) {
	config := Config{path: "/etc/mongodb-backup-agent.conf"}
	err := config.LoadProperties()
	return config, err
}

func (config *Config) GetProperty(name string) string {
	return config.properties[name]
}

func (config *Config) LoadProperties() error {
	config.properties = make(map[string]string)

	file, err := os.Open(config.path)
	if err != nil {
		return &ConfigError{fmt.Sprint(err), err}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), "=")
		config.properties[tokens[0]] = tokens[1]
	}

	if err := scanner.Err(); err != nil {
		return &ConfigError{fmt.Sprint(err), err}
	}
	return nil
}
