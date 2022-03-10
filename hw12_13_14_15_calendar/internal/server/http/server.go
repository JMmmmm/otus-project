package internalhttp

import (
	"context"
	"fmt"
	"github.com/fixme_my_friend/hw12_13_14_15_calendar/domain"
	"github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/logger"
	"net/http"
	"time"
)

type Server struct {
	logger logger.Logger
	app    Application
	server *http.Server
}

type Application interface {
	CreateEvent(ctx context.Context, id, title string) error
	GetEvents(userId int) ([]domain.CalendarEventEntity, error)
	InsertEvent(entities []domain.CalendarEvent) error
	UpdateEvent(entity domain.CalendarEvent) error
	DeleteEvent(userId int) error
}

func NewServer(logger logger.Logger, app Application, addr string) *Server {
	server := Server{
		logger: logger,
		app:    app,
	}
	handler := http.HandlerFunc(server.Handle)
	server.server = &http.Server{Addr: addr, Handler: handler}

	return &server
}

func (s *Server) Start(ctx context.Context) error {
	go func() {
		err := s.server.ListenAndServe()
		s.logger.Error(fmt.Sprintf("stop listening service: %v", err))
	}()

	<-ctx.Done()
	return fmt.Errorf("server stop listening and serve")
}

func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *Server) Handle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	testEvent := &domain.CalendarEvent{
		UserId:               2,
		Title:                "test test",
		DateTimeEvent:        time.Time{},
		DurationEvent:        "3 days 04:05:06",
		Description:          "description",
		NotificationInterval: "3 days 04:05:06",
	}

	testEvent2 := &domain.CalendarEvent{
		UserId:               2,
		Title:                "test2222",
		DateTimeEvent:        time.Time{},
		DurationEvent:        "3 days 04:05:06",
		Description:          "description",
		NotificationInterval: "3 days 04:05:06",
	}

	err := s.app.InsertEvent([]domain.CalendarEvent{*testEvent})
	s.logger.Info(fmt.Sprintf("insert events error: %v", err))

	eventsResult, err := s.app.GetEvents(1)
	s.logger.Info(fmt.Sprintf("get events: %v and error: %v", eventsResult, err))

	err = s.app.UpdateEvent(*testEvent2)
	s.logger.Info(fmt.Sprintf("update events error: %v", err))

	err = s.app.DeleteEvent(2)
	s.logger.Info(fmt.Sprintf("delete events error: %v", err))

	msgID := ""
	w.Header().Add("msgId", msgID)
	w.Write([]byte("Hello, world"))
	s.logger.Warn(fmt.Sprintf("%s %v %s %s %v %s", r.RemoteAddr, start, r.Method, r.URL.Path, time.Since(start), r.UserAgent()))
}
