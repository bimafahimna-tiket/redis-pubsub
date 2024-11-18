package controller

import (
	"net/http"
	"poc-redis-pubsub/internal/application/service"
	"poc-redis-pubsub/internal/domain/dto"

	"github.com/labstack/echo/v4"
)

type MessageController struct {
	messageService service.MessageService
}

func NewMessageController(messageService service.MessageService) *MessageController {
	return &MessageController{
		messageService: messageService,
	}
}

func (c *MessageController) SendMessage(ec echo.Context) error {
	var req dto.MessageRequest
	ctx := ec.Request().Context()
	err := ec.Bind(&req)
	if err != nil {
		return err
	}

	res, err := c.messageService.SendMessage(ctx, req)
	if err != nil {
		return err
	}
	return ec.String(http.StatusOK, res)
}

func (c *MessageController) SendMessagePubSub(ec echo.Context) error {
	ec.Set("entity", "Send-Message")

	var req dto.MessagePubSubRequest
	ctx := ec.Request().Context()
	err := ec.Bind(&req)
	if err != nil {
		return err
	}

	res, err := c.messageService.SendMessagePubSub(ctx, req)
	if err != nil {
		return err
	}
	return ec.String(http.StatusOK, res)
}

func (c *MessageController) GetAllCache(ec echo.Context) error {
	ctx := ec.Request().Context()
	res, err := c.messageService.GetAllCache(ctx)
	if err != nil {
		return err
	}
	return ec.JSON(http.StatusOK, res)
}

func (c *MessageController) UpdateCache(ec echo.Context) error {
	ec.Set("entity", "Update-Cache")

	var req dto.UpdateCacheRequest
	err := ec.Bind(&req)
	if err != nil {
		return err
	}

	ctx := ec.Request().Context()
	res, err := c.messageService.UpdateCache(ctx, req)
	if err != nil {
		return err
	}

	return ec.String(http.StatusOK, res)
}
