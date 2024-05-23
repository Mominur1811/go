package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type DBConfig struct {
	Host     string
	Port     string
	DBName   string
	SSLMode  string
	User     string
	Password string
}

func ReadDBConfigFromFile(filename string) (DBConfig, error) {
	file, err := os.Open(filename)
	if err != nil {
		return DBConfig{}, err
	}
	defer file.Close()

	config := DBConfig{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "=")
		if len(fields) != 2 {
			continue
		}
		key := strings.TrimSpace(fields[0])
		value := strings.TrimSpace(fields[1])

		fmt.Println(key, value)

		switch key {
		case "host":
			config.Host = value
		case "port":
			config.Port = value
		case "dbName":
			config.DBName = value
		case "sslMode":
			config.SSLMode = value
		case "user":
			config.User = value
		case "password":
			config.Password = value
		}
	}

	if err := scanner.Err(); err != nil {
		return DBConfig{}, err
	}

	fmt.Println(config)

	return config, nil
}
