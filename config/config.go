package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	Token        string
	BotPrefix    string
	Server       string
	ClientID     string
	ClientSecret string

	config *configStruct
)

type configStruct struct {
	Token        string `json:"Token"`
	BotPrefix    string `json:"BotPrefix"`
	Server       string `json:"Server"`
	ClientID     string `json:"ClientID"`
	ClientSecret string `json:"ClientSecret"`
}

func ReadConfig() {
	fmt.Println("Reading from config file...")

	file, err := ioutil.ReadFile("./config.json")

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(file))

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err.Error())
	}

	Token = config.Token
	BotPrefix = config.BotPrefix
	Server = config.Server
	ClientID = config.ClientID
	ClientSecret = config.ClientSecret

	return
}
