package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/jukedec/content_server/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
