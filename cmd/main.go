package main

import (
	v1 "project/api/v1"
	"project/core"
	"project/docs"
	"project/infra/config"
	"sync"
)

var wg sync.WaitGroup

// @title        Golang Project Template
// @version      1.0
// @description  API в рамках курса РПМ
//
// @securityDefinitions.apikey  BearerAuth
// @in                          header
// @name                        Authorization
// @description                 Type “Bearer {token}”
func main() {
	server := core.NewHttpServer()
	di := core.NewDi()
	_ = config.Make()
	docs.SwaggerInfo.Host = "127.0.0.1:4114"

	handlers := &v1.Handlers{
		UserHandler: di.UserHandler,
		AuthHandler: di.AuthHandler,
	}
	v1.SetupRoutes(server.App(), handlers)

	wg.Add(1)
	go func() {
		defer wg.Done()
		server.Start()
	}()
	wg.Wait()
}
