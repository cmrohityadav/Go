package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	ServerIp    string      `json:"serverip"`
	Port        int         `json:"port"`
	Bhavcopyurl Bhavcopyurl `json:"bhavcopyurl"`
	Priceband   Priceband   `json:"priceband"`
	Database    DBConfig    `json:"database"`
}

type DBConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
	SSLMode  string `json:"sslmode"`
}

type Bhavcopyurl struct {
	NSE ExchangeDetail `json:"nse"`
	BSE ExchangeDetail `json:"bse"`
}

type Priceband struct {
	NSE ExchangeDetail `json:"nse"`
	BSE ExchangeDetail `json:"bse"`
}

type ExchangeDetail struct {
	URL  string  `json:"url"`
	Time string `json:"time"`
}

func (c *Config) LoadConfig(path string) error {
	pOsFile,err:=os.Open(path);
	

	if err!=nil{
		return fmt.Errorf("error while opening config file: %v", err);
	}

	defer pOsFile.Close();

	json.NewDecoder(pOsFile).Decode(&c);

	return nil;
}
