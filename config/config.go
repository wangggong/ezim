package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
	"qiniupkg.com/x/log.v7"
)

var Config configSt

// LoadConfig loads the config file.
func LoadConfig(filename string) {
	log.Infof("debug: filename = %v", filename)

	s, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("cannot read config file: %v", err)
	}
	if err = yaml.Unmarshal(s, &Config); err != nil {
		log.Fatalf("cannot import config file: %v", err)
	}
}
