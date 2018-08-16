package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Configuration stores the app config
type Configuration struct {
	Database struct {
		Host     string
		Port     int
		User     string
		DBName   string
		Password string
		SSLMode  string
	}
	Ethereum struct {
		ProviderURI string
	}
	Redis struct {
		Path string
	}
}

// Read will load the config file into the Configuration object
func Read(fileName string) (*Configuration, error) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	c := &Configuration{}
	if err = json.Unmarshal(file, &c); err != nil {
		return nil, err
	}

	return c, nil
}

// DBString formats the values to be used when connecting to the db
func (c Configuration) DBString() string {
	return fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=%v",
		c.Database.Host,
		c.Database.Port,
		c.Database.User,
		c.Database.DBName,
		c.Database.Password,
		c.Database.SSLMode,
	)
}
