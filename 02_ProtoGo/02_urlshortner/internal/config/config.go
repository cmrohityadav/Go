package config

import (
	"encoding/json"
	"os"
)

type Config struct{
	Port int `json:"port"`
	ServerIp string `json:"serverIp"`
}
 

func LoadConfig() *Config{

	pJsonfile,err:=os.Open("./config/config.json");
	if err!=nil{
		panic(err);
	}
	defer pJsonfile.Close();

	var cfg Config;
	err=json.NewDecoder(pJsonfile).Decode(&cfg);
	if err!=nil{
		panic(err);
	}

	return &cfg;
}