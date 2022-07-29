package main

import (
	"log"

	"st-booker/internal/server"
	"st-booker/internal/service"
	"st-booker/internal/storage"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	lg, err := zap.NewDevelopment(zap.Development())
	if err != nil {
		log.Fatalln("failed to init logger")
	}

	st := storage.NewClient()

	srv := service.NewBooking(st, lg)

	r := gin.Default()
	s := server.NewServer(r, srv, lg)
	s.InitRoutes()

	if err := r.Run(":8080"); err != nil { //TODO: move port to config
		lg.Error(err.Error())
	}
}
