package router

import (
	"echo-playground/conf"
	ctl "echo-playground/controller"

	"github.com/labstack/echo/v4"
)

type Router struct {
	cfg     *conf.Config
	ct      *ctl.Controller
	hHealth *ctl.Health
}

func NewRouter(cf *conf.Config, ct *ctl.Controller) (*Router, error) {
	r := &Router{
		cfg:     cf,
		ct:      ct,
		hHealth: ct.GetHealthHandler(),
	}
	return r, nil
}

func (p *Router) Idx() *echo.Echo {
	e := echo.New()
	// TODO: Add logger, CORS, etc...

	// e.GET("/health", p.ct.GetHealthHandler().Check)
	e.GET("/health", p.hHealth.Check)

	return e
}
