package httpserver

import (
	"auth/internal/config"
	"auth/internal/db"
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type MyHTTPServerApp struct {
	Config config.Config
	DB     *pgxpool.Pool
	Server *http.Server
}

func (myServer *MyHTTPServerApp)NewSetupServer()(error){
	cfg,err:=config.LoadConfig()
	if err!=nil{
		return	fmt.Errorf("Unable to Start Server Config File are not proper: %s",err.Error())
	}

	dbPool,err:=db.ConnectToDatabase(cfg);
	if err!=nil{
		return fmt.Errorf("Unable to connect DB Check .env file :%s",err.Error())
	}
	
	router := NewRouter(dbPool,cfg)

	myServer.Server = &http.Server{
		Addr:    ":" + cfg.SERVER_PORT,
		Handler: router,
	}

	myServer.Config=cfg
	myServer.DB=dbPool

	return nil
}

func (app *MyHTTPServerApp) Start() error {
	return app.Server.ListenAndServe()
}

