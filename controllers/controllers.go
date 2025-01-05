package controllers

import (
	"fmt"
	"net/http"

	"github.com/jocsakesley/server-api/infra/http/server"
)

func GetName(c server.IContext) error {
	name := c.Param("name")
	return c.JSON(http.StatusOK, map[string]string{"message": fmt.Sprintf("Hello, %s", name)})
}

func PostValues(c server.IContext) error {
	body := map[string]string{}
	if err := c.BodyParser(&body); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, body)
}
