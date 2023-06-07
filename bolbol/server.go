package bolbol

import (
	"context"
	"github.com/labstack/echo/v4"
	"push_notification/entity"
	"strconv"
	"time"
)

type Server struct {
	bolbol Bolbol
}

type NotifyRequest struct {
	UserID            int                               `json:"user_id"`
	UnreadMessage     *entity.UnreadMessageNotification `json:"unread_message"`
	UnreadWorkRequest *entity.UnreadWorkRequest         `json:"unread_work_request"`
}

func (s *Server) listen(c echo.Context) error {
	clientId, _ := strconv.Atoi(c.Param("id"))
	notification, err := s.bolbol.getNotification(c.Request().Context(), clientId)
	if err != nil {
		return err
	}

	return c.JSON(200, notification)
}

func (n *NotifyRequest) Notification() entity.Notification {
	if n.UnreadMessage != nil {
		return n.UnreadMessage
	}
	if n.UnreadWorkRequest != nil {
		return n.UnreadWorkRequest
	}
	panic("bad notification")

}

func (s *Server) notify(c echo.Context) error {
	var request NotifyRequest
	if err := c.Bind(&request); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := s.bolbol.Notify(ctx, request.UserID, request.Notification()); err != nil {
		return err
	}
	return c.String(201, "notification created")
}
