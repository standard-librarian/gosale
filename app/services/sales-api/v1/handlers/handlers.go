package handlers

import (
	"github.com/standard-librarian/gosale/app/services/sales-api/v1/handlers/hackgrp"
	v1 "github.com/standard-librarian/gosale/business/web/v1"
	"github.com/standard-librarian/gosale/foundation/web"
)

type Routes struct{}

// Add implements the RouterAdder interface.
func (Routes) Add(app *web.App, cfg v1.APIMuxConfig) {
	hackgrp.Routes(app)
}
