package config

import (
	"encoding/json"
	"io/ioutil"
)

// Config holds all the required information for server.
type Config struct {
	Proxy struct {
		DNS struct {
			Enabled  bool   `json:"enabled"`
			Listener string `json:"listener"`
		} `json:"dns"`
		SSL struct {
			Enabled  bool   `json:"enabled"`
			Listener string `json:"listener"`
			Key      string `json:"key"`
			Cert     string `json:"cert"`
		} `json:"ssl"`
		HTTP struct {
			Enabled  bool   `json:"enabled"`
			Listener string `json:"listener"`
		} `json:"http"`
	} `json:"proxy"`
	Admin struct {
		Listener string `json:"listener"`
		Key      string `json:"key"`
		Cert     string `json:"cert"`
	} `json:"admin"`
	JWTKey     string `json:"jwt_key"`
	BoltDBFile string `json:"bolt_db_file"`
}

// New parses JSON from the file provided by filename into a Config struct.
func New(filename string) (*Config, error) {
	config := &Config{}
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return config, err
	}

	if err := json.Unmarshal(file, &config); err != nil {
		return config, err
	}

	return config, nil
}
