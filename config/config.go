package config

import (
	"log"
	"github.com/BurntSushi/toml"
)
// Info from config file
type Config struct {
	From_name   	string
	From_email     	string
	To_name  	string
	Subject    	string
	Content     	string
	Sg_url  	string
	Sg_route	string
}

// Reads info from config file
func ReadConfig() Config {
	var config Config
	if _, err := toml.DecodeFile("/home/ec2-user/work/src/github.com/noamelya/mobilize/config/config.toml", &config); err != nil {
		log.Fatal("problem fetching conf file")
	}

	return config
}
