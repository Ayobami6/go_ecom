package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Ayobami6/go_ecom/config"
	"github.com/Ayobami6/go_ecom/db"
	myscf "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	db, err := db.NewSQLStorage(myscf.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPasswd,
		DBName:               config.Envs.DBName,
		Addr:                 config.Envs.DBAddress,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}
	driver, err := mysql.WithInstance(db, &mysql.Config{}) 
	if err!= nil {
        log.Fatal(err)
    }
	m, merr := migrate.NewWithDatabaseInstance("file://cmd/migrate/migrations", "mysql",
		driver,
	)
	if merr!= nil {
        log.Fatal(err)
    }
	cmd := os.Args[(len(os.Args) - 1)]
	fmt.Println(cmd)
	if cmd == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			panic(err)
		}
	}
	if cmd == "down" {
        if err := m.Down(); err!= nil && err!= migrate.ErrNoChange {
            panic(err)
        }
    }

}
