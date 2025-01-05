package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/jocsakesley/server-api/infra/http/server"
)

type serverGin struct {
	engine *gin.Engine
}

func NewGinServer() server.IServer {
	return &serverGin{
		engine: gin.Default(),
	}
}

func handlerGinAdapter(handler interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := ginContext{c: c}
		h := handler.(func(ctx server.IContext) error)
		h(ctx)
	}
}

func (s *serverGin) Get(path string, handler interface{}) {
	s.engine.GET(path, handlerGinAdapter(handler))
}

func (s *serverGin) Post(path string, handler interface{}) {
	s.engine.POST(path, handlerGinAdapter(handler))
}

func (s *serverGin) Run(port string) {
	s.engine.Run(port)
}
