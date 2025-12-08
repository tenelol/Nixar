// apps/portfolio/handlers.go
package portfolio

import (
	"net/http"

	"github.com/tenelol/nixar/framework"
)

type Project struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
    URL         string `json:"url"`
}


var projects = []Project{
    {
        ID:          4,
        Name:        "nixosをサーバー化した話",
        Description: "my first qiita journal",
	URL:         "https://qiita.com/tener/items/940198e079b20385cfab",
    },

    {
        ID:          3,
        Name:        "nixos-dotfiles",
        Description: "my nix serverized configuration",
	URL:         "https://github.com/tenelol/nixos-dotfiles",
    },

    {
        ID:          2,
        Name:        "My Framework",
        Description: "my portfolio configuration with self-made framework",
	URL:         "https://github.com/tenelol/nixar", 
    },

        {
        ID:          1,
        Name:        "linux-notes",
        Description: "my first notes to learn linux",
	URL:         "https://github.com/tenelol/linux-notes", 
    },
    // 必要なだけ続ける
}


// GET /api/projects
func ProjectsAPI(ctx *framework.Context) {
    ctx.JSON(http.StatusOK, projects)
}
