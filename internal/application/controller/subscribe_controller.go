package controller

import (
	"net/http"
	"poc-redis-pubsub/internal/application/service"
	"poc-redis-pubsub/internal/domain/dto"

	"github.com/labstack/echo/v4"
)

type SubscribeController struct {
	subscribeService service.SubscribeService
}

func NewSubscribeController(subscribeService service.SubscribeService) *SubscribeController {
	return &SubscribeController{
		subscribeService: subscribeService,
	}
}

func (c *SubscribeController) SubscribeToChannel(ec echo.Context) error {
	var req dto.SubscribeToChannelRequest

	ctx := ec.Request().Context()
	err := ec.Bind(&req)
	if err != nil {
		return err
	}
	res, err := c.subscribeService.SubscribeTo(ctx, req.Channel)
	if err != nil {
		return ec.String(http.StatusBadRequest, err.Error())
	}
	return ec.String(http.StatusOK, res)
}

func (c *SubscribeController) UnsubscribeToChannel(ec echo.Context) error {
	var req dto.UnsubscribeToChannelRequest

	ctx := ec.Request().Context()
	err := ec.Bind(&req)
	if err != nil {
		return err
	}
	res, err := c.subscribeService.UnsubscribeTo(ctx, req.Channel)
	if err != nil {
		return ec.String(http.StatusBadRequest, err.Error())
	}
	return ec.String(http.StatusOK, res)
}
