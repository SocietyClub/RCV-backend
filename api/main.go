package main

import (
	"context"
	"fmt"
	app "go-service/controller"
	"net/http"
	"strconv"

	"github.com/core-go/config"
	"github.com/core-go/log"
	mid "github.com/core-go/log/middleware"
	"github.com/gorilla/mux"
)

func main() {

	// Loading Configs from Service Account
	fmt.Println("Loading Configs")
	var conf app.Root
	error1 := config.Load(&conf, "configs/config")
	if error1 != nil {
		panic(error1)
	}

	// Setting up Mux Route and using configs
	r := mux.NewRouter()

	log.Initialize(conf.Log)
	r.Use(mid.BuildContext)
	logger := mid.NewStructuredLogger()
	if log.IsInfoEnable() {
		r.Use(mid.Logger(conf.MiddleWare, log.InfoFields, logger))
	}
	r.Use(mid.Recover(log.ErrorMsg))

	error2 := app.Route(r, context.Background(), conf)
	if error2 != nil {
		panic(error2)
	}

	// Start Server
	fmt.Println("RCV Backend Google Cloud Server Started")
	server := ""
	if conf.Server.Port > 0 {
		server = ":" + strconv.FormatInt(conf.Server.Port, 10)
	}
	if error3 := http.ListenAndServe(server, r); error3 != nil {
		fmt.Println(error3.Error())
	}
}
