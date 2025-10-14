package core

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"runtime"
	"strings"
	"time"
)

type HttpServer struct {
	app *fiber.App
}

type Server interface {
	Start()
	App() *fiber.App
}

func NewHttpServer() Server {
	app := fiber.New(fiber.Config{
		BodyLimit:         1024 * 1024,
		AppName:           "Golang Template",
		StreamRequestBody: true,
	})

	var methods = []string{fiber.MethodGet, fiber.MethodDelete, fiber.MethodPost, fiber.MethodPut}
	var headers = []string{fiber.HeaderAccept, fiber.HeaderContentType, fiber.HeaderAccept, fiber.HeaderAuthorization}

	corsConfig := cors.New(cors.Config{
		AllowOrigins:     strings.Join([]string{"http://localhost:4114"}, ", "),
		AllowMethods:     strings.Join(methods, ", "),
		AllowHeaders:     strings.Join(headers, ", "),
		AllowCredentials: true,
		MaxAge:           300,
	})

	app.Use(corsConfig)
	app.Use(recover2.New())
	app.Use(logger.New(logger.Config{
		Format:       "${time} | ${status} | ${latency} | ${ip} | ${method} | ${path} | ${error}\n",
		TimeFormat:   "15:04:05",
		TimeZone:     "Europe/Moscow",
		TimeInterval: 500 * time.Millisecond,
	}))

	return &HttpServer{
		app: app,
	}
}

func (s *HttpServer) App() *fiber.App {
	return s.app
}

func (s *HttpServer) Start() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	if err := s.app.Listen(":" + "4114"); err != nil {
		panic(err)
	}
}
