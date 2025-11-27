package main

import (
	"database/sql"
	"log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {
	e := echo.New();
	e.Use(middleware.Logger())

	//creating DSN;
	dsn:="host=localhost port=5432 user=postgres password=123456 dbname=echol sslmode=disable"

	db,err:=sql.Open("postgres",dsn);

	if err!=nil{
		log.Fatal("failed to connect the database...",err);
	}

	defer db.Close();

	if err:=db.Ping(); err!=nil{
		log.Fatal("Database ping failed---",err);
	}

	createTable:=`
	CREATE TABLE IF NOT EXISTS users(
		id serial primary key,
		name text,
		email text unique,
		age int
	);
	`
	if _,err:=db.Exec(createTable);err!=nil{
		log.Fatal("failed to create the table");
	}

	e.Logger.Fatal(e.Start(":8000"));

}
