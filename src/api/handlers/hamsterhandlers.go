package handlers

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type Hamster struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func AddHamster(c echo.Context) error {
	hamster := Hamster{}

	err := c.Bind(&hamster)
	if err != nil {
		log.Printf("Failed processing addHamster request: %s\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("this is your hamster: %#v", hamster)
	return c.String(http.StatusOK, "we got your hamster!")
}
