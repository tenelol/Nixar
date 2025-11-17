// apps/portfolio/handlers.go
package portfolio

import (
	"net/http"

	"mywebfw/framework"
)

type Project struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
    URL         string `json:"url"`
}


var projects = []Project{
    {
        ID:          1,
        Name:        "dotfiles",
        Description: ".config",
	URL:         "https://github.com/tener-kiwi",
    },
    {
        ID:          2,
        Name:        "My Framework",
        Description: "自作 Go Web フレームワーク",
	URL:         "https://github.com/tener-kiwi", // ここは適当に
    },
    // 必要なだけ続ける
}


// GET /api/projects
func ProjectsAPI(ctx *framework.Context) {
	framework.JSON(ctx, http.StatusOK, projects)
}

