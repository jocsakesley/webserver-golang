package routes

import (
	"github.com/jocsakesley/server-api/controllers"
	"github.com/jocsakesley/server-api/infra/http/server"
)

func SetupRoutes(server server.IServer) server.IServer {
	server.Get("/:name", controllers.GetName)
	server.Post("/", controllers.PostValues)
	return server
}
