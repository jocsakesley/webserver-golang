package main

import (
	r "github.com/jocsakesley/server-api/infra/http/routes"
	g "github.com/jocsakesley/server-api/infra/http/server/gin"
	//f "github.com/jocsakesley/server-api/infra/http/server/fiber"
)

func main() {
	//server := f.NewFiberServer()
	server := g.NewGinServer()
	router := r.SetupRoutes(server)

	router.Run(":8080")
}
