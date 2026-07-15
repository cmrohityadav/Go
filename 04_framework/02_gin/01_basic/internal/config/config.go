package config

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)
type Config struct {
    DatabaseURL string
    ServerPort  string
}

func LoadConfig()(Config,error){
	if err:=godotenv.Load() ; err!=nil{
		return Config{}, fmt.Errorf("failed to load .env: %w", err)
	}

	sPgsqlUrl,err:=getEnv("POSTGRESQL_URL")
	if err!=nil{
		return Config{},fmt.Errorf("%s",err.Error())
	}
	
	sServerPort,err:=getEnv("SERVER_PORT")
	if err!=nil{
		return Config{},fmt.Errorf("%s",err.Error())
	}

	return Config{
		DatabaseURL: sPgsqlUrl,
		ServerPort: sServerPort,
	}, nil

}

func getEnv(key string)(string,error){
	val:=os.Getenv(key)

	if val==""{
		return "", fmt.Errorf("missing environment variable: %s", key)
	}

	return val,nil
}