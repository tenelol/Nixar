package web

import (
	"net/http"

	"nixar/framework"
)

func Index(c *framework.Context) {
	c.JSON(http.StatusOK, map[string]any{
		"message": "Hello from Nixar!",
	})
}

