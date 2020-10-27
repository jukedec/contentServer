package actions

import (
    "net/http"
    
	"github.com/gobuffalo/buffalo"
)

// MyGuyIndex default implementation.
func MyGuyIndex(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("my_guy/index.html"))
}

