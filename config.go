/*

configuration management

*/

package main

import (
	"encoding/json"
	"log"
	"os"
)

// Configuration struct with all variables imported from config.json
type Configuration struct {
	ServerPort int `json:"server_port"`

	RootFolder string `json:"root_folder"`
}

var conf Configuration

// ReadConfig reads 'config.json' and fills Configuration struct
func ReadConfig() {
	configFile := "webservermax-config.json"
	log.Printf("Reading configuration from '%s'", configFile)
	var err error
	conf, err = loadConfig(configFile)
	if err == nil {
		log.Printf("%+v\n", conf)
	} else {
		log.Println(err)
	}
}

func saveConfig(c Configuration, filename string) error {
	bytes, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, bytes, 0644)
}

func loadConfig(filename string) (Configuration, error) {

	bytes, err := os.ReadFile(filename)
	if err != nil {
		return Configuration{}, err
	}

	var c Configuration
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return Configuration{}, err
	}

	return c, nil
}

/*
func main() {
	configuration, err := loadConfig("config.json")
	if err == nil {
		fmt.Printf("%+v\n", configuration)
	} else {
		fmt.Println(err)
	}
}
*/
