package db

import (
	"auth/internal/config"
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func MakeDBConnectionStringForPGSQL(config config.Config) string{
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",config.DB_USER_NAME,config.DB_PASSWORD,config.DB_HOST_IP,config.DB_PORT,config.DB_NAME,config.DB_MODE)
}

func ConnectToDatabase(cfg config.Config)(*pgxpool.Pool,error){
	ctx,cancle:=context.WithTimeout(context.Background(),time.Second*10)
	defer cancle()

	pool,err:=pgxpool.New(ctx,MakeDBConnectionStringForPGSQL(cfg))
	if err!=nil{
		pool.Close()
		return nil,err
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}

	return pool,nil

}
