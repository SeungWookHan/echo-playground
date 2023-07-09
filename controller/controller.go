package controller

import (
	"echo-playground/conf"
	"echo-playground/protocol"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	cfg           *conf.Config
	healthHandler *Health
}

func NewCTL(cf *conf.Config) (*Controller, error) {
	r := &Controller{
		cfg: cf,
	}

	r.healthHandler = NewHeartbeat(r)
	return r, nil
}

func (p *Controller) RespOK(c echo.Context, resp interface{}) error {
	c.JSON(http.StatusOK, resp)
	return nil
}

func (p *Controller) RespError(c echo.Context, body interface{}, status int, err ...interface{}) {
	// bytes, _ := json.Marshal(body)

	// TODO: log

	c.JSON(status, protocol.NewRespHeader(protocol.Failed, joinMsg(c.Request().RequestURI, err)))
}

func (p *Controller) GetHealthHandler() *Health {
	return p.healthHandler
}
