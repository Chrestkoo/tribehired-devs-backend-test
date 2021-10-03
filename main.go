package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/chrestkoo/project/tribehired-devs-backend-test/pkg/logging"
	"github.com/chrestkoo/project/tribehired-devs-backend-test/pkg/setting"
	"github.com/chrestkoo/project/tribehired-devs-backend-test/routers"
)

func init() {
	setting.Setup()
	logging.Setup()
}

// @title Golang Gin API
// @version 1.0
// @description An example of gin
// @license.name MIT
func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf("%v:%d", setting.ServerSetting.Domain, setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
