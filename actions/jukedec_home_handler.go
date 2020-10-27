package actions

import (
    "net/http"
    
	"github.com/gobuffalo/buffalo"
)

// JukedecHomeHandlerIndex default implementation.
func JukedecHomeHandlerIndex(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("jukedec_home_handler/index.html"))
}



// JukedecHomeHandlerJukedecHomeHandler default implementation.
func JukedecHomeHandlerJukedecHomeHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("jukedec_home_handler/jukedec_home_handler.html"))
}


// JukedecHomeHandlerSupermyhomehandler default implementation.
func JukedecHomeHandlerSupermyhomehandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("jukedec_home_handler/supermyhomehandler.html"))
}


// JukedecHomeHandlerLs default implementation.
func JukedecHomeHandlerLs(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("jukedec_home_handler/ls.html"))
}

