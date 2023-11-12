package main

import (
	"database/sql"
	"github.com/Wintersunner/xplor/api"
	db "github.com/Wintersunner/xplor/db/sqlc"
	"github.com/Wintersunner/xplor/util"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattes/migrate/source/file"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config file", err)
		return
	}

	migrateDatabase(config)

	conn, err := sql.Open(config.DBDriver, config.DBSource())

	if err != nil {
		log.Fatal("cannot connect to main database: ", err)
	}

	store := db.NewStore(conn)

	server := api.NewServer(store, config)

	err = server.Start()

	if err != nil {
		log.Fatal("cannot start server", err)
	}

}

func migrateDatabase(config util.Config) {
	migration, err := migrate.New("file://db/migration", config.MigrationSource())

	if err != nil {
		log.Fatal("cannot create new migration instance: ", err)
	}

	err = migration.Up()

	if err != nil && err != migrate.ErrNoChange {
		log.Fatal("failed to run database migrations: ", err)
	}

	log.Println("database migrated successfully!")
}
