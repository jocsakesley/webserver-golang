package fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jocsakesley/server-api/infra/http/server"
)

type serverFiber struct {
	app *fiber.App
}

func NewFiberServer() server.IServer {
	return &serverFiber{
		app: fiber.New(),
	}
}

func handlerFiberAdapter(handler interface{}) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		ctx := fiberContext{c: c}
		h := handler.(func(ctx server.IContext) error)
		return h(ctx)
	}
}

func (s *serverFiber) Get(path string, handler interface{}) {
	s.app.Get(path, handlerFiberAdapter(handler))
}

func (s *serverFiber) Post(path string, handler interface{}) {
	s.app.Post(path, handlerFiberAdapter(handler))
}

func (s *serverFiber) Run(port string) {
	s.app.Listen(port)
}
