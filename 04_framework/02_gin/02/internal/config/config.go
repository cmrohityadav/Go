package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_USER_NAME string
	DB_PASSWORD string
	DB_HOST_IP string
	DB_PORT string
	DB_NAME string
	DB_MODE string
	SERVER_PORT string
	JWT_SECRET string
}

func LoadConfig() (Config, error) {
	err:=godotenv.Load(".env")
	if err!=nil{
		return Config{},err;
	}

	dbUserName,err:=getEnv("DB_USER_NAME")
	if err!=nil{
		return Config{},err
	}
	dbPassword,err:=getEnv("DB_PASSWORD")
	if err!=nil{
		return Config{},err
	}

	dbHostIp,err:=getEnv("DB_HOST_IP")
	if err!=nil{
		return Config{},err
	}

	dbPort,err:=getEnv("DB_PORT")
	if err!=nil{
		return Config{},err
	}
	
	dbname,err:=getEnv("DB_NAME")
	if err!=nil{
		return Config{},err
	}
	
	dbMode,err:=getEnv("DB_MODE")
	if err!=nil{
		return Config{},err
	}

	serverPort,err:=getEnv("SERVER_PORT")
	if err!=nil{
		return Config{},err
	}

	jwtSecret,err:=getEnv("JWT_SECRET")
	if err!=nil{
		return Config{},err
	}

	

	return Config{
		DB_USER_NAME: dbUserName,
		DB_PASSWORD: dbPassword,
		DB_NAME: dbname,
		DB_PORT: dbPort,
		DB_HOST_IP: dbHostIp,
		DB_MODE: dbMode,
		SERVER_PORT: serverPort,
		JWT_SECRET:jwtSecret ,
	},nil
	
}

func getEnv(key string)(string,error){
	value:=os.Getenv(key)
	if value==""{
		return "",fmt.Errorf("Missing Environment Variable %s ",key)
	}

	return strings.TrimSpace(value),nil;
}
