package simple

import (
	"net/http"

	"github.com/tenelol/nixar/framework"
)

type HelloResponse struct {
	Message string `json:"message"`
}

// GET /api/hello
func HelloAPI(ctx *framework.Context) {
	ctx.JSON(http.StatusOK, HelloResponse{Message: "hello from nixar"})
}
