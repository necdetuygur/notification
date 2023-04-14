package services

import (
	"net/http"
	"notification/functions"

	"github.com/labstack/echo/v4"
)

func Send(c echo.Context) error {
	title := c.QueryParam("title")
	body := c.QueryParam("body")
	link := c.QueryParam("link")
	topic := c.QueryParam("topic")
	response := functions.NotificationSend(title, body, link, topic)
	return c.String(http.StatusOK, response)
}
