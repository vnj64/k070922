package cmd

import (
	"github.com/swaggo/swag/example/basic/docs"
	v1 "project/api/v1"
	"project/core"
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
	cfg := config.Make()
	docs.SwaggerInfo.Host = cfg.FullHttpHost() + cfg.HttpPort() + "/"

	handlers := &v1.Handlers{
		UserHandler: di.UserHandler,
	}
	v1.SetupRoutes(server.App(), handlers)
	wg.Add(1)
	go func() {
		defer wg.Done()
		server.Start()
	}()
	wg.Wait()
}
