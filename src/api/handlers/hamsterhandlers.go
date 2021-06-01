package handlers

import (
	"log"
	"myapp/src/api/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddHamster(c echo.Context) error {
	hamster := models.Hamster{}

	err := c.Bind(&hamster)
	if err != nil {
		log.Printf("Failed processing addHamster request: %s\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("this is your hamster: %#v", hamster)
	return c.String(http.StatusOK, "we got your hamster!")
}
