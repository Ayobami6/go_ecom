package main

import (
	"database/sql"
	"log"

	"github.com/Ayobami6/go_ecom/cmd/api"
	"github.com/Ayobami6/go_ecom/config"
	"github.com/Ayobami6/go_ecom/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	DB, err := db.NewSQLStorage(mysql.Config{
		User:   config.Envs.DBUser,
        Passwd: config.Envs.DBPasswd,
        DBName:     config.Envs.DBName,
		Addr: config.Envs.DBAddress,
		Net: "tcp",
		AllowNativePasswords: true,
		ParseTime: true,
	})
	if err!= nil {
        log.Fatal(err)
    }
	initStorage(DB)
	server := api.NewAPIServer("localhost:5000", DB)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err!= nil {
        log.Fatal(err)
    }
	log.Println("Database Connected Successfully")
}