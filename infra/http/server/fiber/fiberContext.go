package fiber

import "github.com/gofiber/fiber/v2"

type fiberContext struct {
	c *fiber.Ctx
}

func (f fiberContext) Param(key string) string {
	return f.c.Params(key)
}

func (f fiberContext) JSON(code int, object any) error {
	f.c.Status(code)
	return f.c.JSON(object)
}

func (f fiberContext) BodyParser(out interface{}) error {
	return f.c.BodyParser(out)
}
