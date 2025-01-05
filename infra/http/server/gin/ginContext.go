package gin

import "github.com/gin-gonic/gin"

type ginContext struct {
	c *gin.Context
}

func (g ginContext) Param(key string) string {
	return g.c.Param(key)
}

func (g ginContext) JSON(code int, object any) error {
	g.c.JSON(code, object)
	return nil
}

func (g ginContext) BodyParser(out interface{}) error {
	return g.c.BindJSON(out)
}
