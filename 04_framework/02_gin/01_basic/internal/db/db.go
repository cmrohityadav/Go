package db

import (
	"context"
	"ginlearn/internal/config"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)


func Connect(cfg config.Config) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}

	return pool, nil
}

func Disconnect(pool *pgxpool.Pool) {
	if pool != nil {
		pool.Close()
	}
}


func Connect_for_learning(cfg config.Config)(*pgx.Conn,error){

	ctx,cancel:=context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()

	conn,err:=pgx.Connect(ctx,cfg.DatabaseURL)

	if err != nil {
		return nil, err
	}

	return conn, nil

}

func DisconnectConn_for_learning(conn *pgx.Conn) error {
	if conn == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return conn.Close(ctx)
}
