package internalhttp

import (
	"context"
	"fmt"
	"net/http"
	"time"

	domain "github.com/fixme_my_friend/hw12_13_14_15_calendar/domain/calendarevent"
	"github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/logger"
)

type Server struct {
	logger logger.Logger
	app    Application
	server *http.Server
}

type Application interface {
	GetEvents(userID int) ([]domain.CalendarEventEntity, error)
	CreateEvent(title string, dateTimeEvent time.Time, DurationEvent string, userID int) error
	UpdateEvent(id int, title string) error
	DeleteEvent(userID int) error
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

	err := s.app.CreateEvent("test test", time.Time{}, "3 days 04:05:06", 2)
	s.logger.Info(fmt.Sprintf("insert events error: %v", err))

	eventsResult, err := s.app.GetEvents(2)
	s.logger.Info(fmt.Sprintf("get events: %v and error: %v", eventsResult, err))

	err = s.app.UpdateEvent(2, "test2222")
	s.logger.Info(fmt.Sprintf("update events error: %v", err))

	eventsResult, err = s.app.GetEvents(2)
	s.logger.Info(fmt.Sprintf("get events: %v and error: %v", eventsResult, err))

	err = s.app.DeleteEvent(2)
	s.logger.Info(fmt.Sprintf("delete events error: %v", err))

	msgID := ""
	w.Header().Add("msgId", msgID)
	w.Write([]byte("Hello, world"))

	message := fmt.Sprintf(
		"%s %v %s %s %v %s",
		r.RemoteAddr, start, r.Method, r.URL.Path, time.Since(start), r.UserAgent())
	s.logger.Warn(message)
}
