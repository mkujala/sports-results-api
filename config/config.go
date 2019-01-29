package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// Config is type for config values
type Config struct {
	Env    string `json:"Env"`
	Port   string `json:"Port"`
	APIURL string `json:"ApiURL"`
	DB     struct {
		Name string `json:"Name"`
		URI  string `json:"URI"`
		User string `json:"User"`
		Pass string `json:"Pass"`
	} `json:"DB"`
}

// Values returns config values
func Values() Config {
	jsonFile, err := os.Open("config/config.json")

	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	var config Config
	json.Unmarshal(byteValue, &config)

	return config
}
