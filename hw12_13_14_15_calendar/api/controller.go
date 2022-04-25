package api

import (
	"context"
	"time"

	"github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/app"
	calendar_event_api "github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/pkg/calendar-event"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Controller struct {
	calendar_event_api.UnimplementedCalendarEventServer
	App app.Application
}

func (c *Controller) GetEvents(ctx context.Context,
	req *calendar_event_api.GetUserEventsRequest) (*calendar_event_api.Events, error) {
	entities, err := c.App.GetEvents(int(req.UserId))
	if err != nil {
		return nil, err
	}

	userEvents := make([]*calendar_event_api.Event, len(entities))
	for key, entity := range entities {
		userEvents[key] = &calendar_event_api.Event{
			Id:     entity.ID,
			UserId: int64(entity.UserID),
			Title:  entity.Title,
		}
	}

	return &calendar_event_api.Events{Content: userEvents}, nil
}

func (c *Controller) CreateEvent(ctx context.Context,
	req *calendar_event_api.CreateEventRequest) (*emptypb.Empty, error) {
	err := c.App.CreateEvent(req.GetTitle(), time.Time{}, req.GetDuration(), int(req.GetUserId()))

	return &empty.Empty{}, err
}

func (c *Controller) UpdateEvent(ctx context.Context,
	req *calendar_event_api.UpdateEventRequest) (*emptypb.Empty, error) {
	err := c.App.UpdateEvent(req.GetId(), req.GetTitle())

	return &empty.Empty{}, err
}

func (c *Controller) DeleteEvent(ctx context.Context,
	req *calendar_event_api.DeleteEventRequest) (*emptypb.Empty, error) {
	err := c.App.DeleteEvent(req.GetId())

	return &empty.Empty{}, err
}
