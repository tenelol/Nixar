package web

import (
	"net/http"

	"github.com/tenelol/nixar/framework"
)

type HelloResponse struct {
	Message string `json:"message"`
}

// GET /api/hello
func HelloAPI(c *framework.Context) {
	c.JSON(http.StatusOK, HelloResponse{Message: "hello from nixar template"})
}
