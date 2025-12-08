package web

import (
	"net/http"

	"github.com/tenelol/nixar/framework"
)

func Index(c *framework.Context) {
	c.JSON(http.StatusOK, map[string]any{
		"message": "Hello from Nixar!",
	})
}

