package controller

import (
	"echo-playground/conf"
	"echo-playground/protocol"

	"github.com/labstack/echo/v4"
)

type Health struct {
	ctl *Controller
	cfg *conf.Config
}

func NewHeartbeat(h *Controller) *Health {
	return &Health{
		ctl: h,
	}
}

// @Summary check
// @Description 서버 상태 체크
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} protocol.RespHeader
// @Router /health/check [get]
func (p *Health) Check(c echo.Context) error {
	return p.ctl.RespOK(c, protocol.NewRespHeader(protocol.Success))
}
