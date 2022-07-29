package main

import (
	"database/sql"
	"fmt"
	"log"

	"st-booker/internal/config"
	"st-booker/internal/server"
	"st-booker/internal/service"
	"st-booker/internal/storage"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"go.uber.org/zap"
)

func main() {
	lg, err := zap.NewDevelopment(zap.Development())
	if err != nil {
		log.Fatalln("failed to init logger")
	}

	cfg, err := config.ReadConfig()
	if err != nil {
		lg.Error(err.Error())
	}

	db, err := sql.Open("postgres", connStr(&cfg))
	if err != nil {
		lg.Error(err.Error())
	}

	err = migrateUp(db)
	if err != nil {
		lg.Error(err.Error())
	}

	st, err := storage.NewClient(db)

	srv := service.NewBooking(st, lg)

	r := gin.Default()
	s := server.NewServer(r, srv, lg)
	s.InitRoutes()

	if err := r.Run(":8080"); err != nil { //TODO: move port to config
		lg.Error(err.Error())
	}
}

func migrateUp(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil {
		return err
	}

	return nil
}

func connStr(c *config.Config) string {
	pConfig := c.DB

	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s ",
		pConfig.Host, pConfig.Port, pConfig.Username, pConfig.DBName, pConfig.Password, pConfig.SSLMode)
}
