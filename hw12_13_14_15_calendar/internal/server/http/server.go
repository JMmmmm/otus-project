package internalhttp

import (
	"context"
	"fmt"
	calendar_event_api "github.com/fixme_my_friend/hw12_13_14_15_calendar/pkg/calendar-event"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/felixge/httpsnoop"
	domain "github.com/fixme_my_friend/hw12_13_14_15_calendar/domain/calendarevent"
	"github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/logger"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
)

type Server struct {
	logger       logger.Logger
	server       cmux.CMux
	httpServer   *http.Server
	grpcServer   *grpc.Server
	httpListener net.Listener
	grpcListener net.Listener
}

type Controller struct {
	calendar_event_api.UnimplementedCalendarEventServer
	app Application
}

func (c *Controller) GetEvents(ctx context.Context, req *calendar_event_api.GetUserEventsRequest) (*calendar_event_api.Events, error) {
	entities, err := c.app.GetEvents(int(req.UserId))
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

type Application interface {
	GetEvents(userID int) ([]domain.CalendarEventEntity, error)
	CreateEvent(title string, dateTimeEvent time.Time, DurationEvent string, userID int) error
	UpdateEvent(id int, title string) error
	DeleteEvent(userID int) error
}

func withLogger(handler http.Handler, logger logger.Logger) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		m := httpsnoop.CaptureMetrics(handler, writer, request)
		logger.Info(fmt.Sprintf("%s-- %v -- http[%d]-- %s -- %s -- %s\n", request.RemoteAddr, time.Now(), m.Code, m.Duration, request.URL.Path, request.UserAgent()))
	})
}

func NewServer(logger logger.Logger, app Application, addr string) *Server {
	grpcSever := grpc.NewServer()

	calendar_event_api.RegisterCalendarEventServer(grpcSever, &Controller{
		app: app,
	})
	mux := runtime.NewServeMux()
	err := calendar_event_api.RegisterCalendarEventHandlerFromEndpoint(context.Background(), mux, "localhost:8081", []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatal(err)
	}
	server := &http.Server{
		Handler: withLogger(mux, logger),
	}
	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	m := cmux.New(l)

	return &Server{
		logger:       logger,
		server:       m,
		httpServer:   server,
		grpcServer:   grpcSever,
		httpListener: m.Match(cmux.HTTP1Fast()),
		grpcListener: m.Match(cmux.HTTP2()),
	}
}

func (s *Server) Start(ctx context.Context) error {
	go func() {
		err := s.httpServer.Serve(s.httpListener)
		s.logger.Error(fmt.Sprintf("stop listening http server: %v", err))
	}()

	go func() {
		err := s.grpcServer.Serve(s.grpcListener)
		s.logger.Error(fmt.Sprintf("stop listening grpc server: %v", err))
	}()

	go func() {
		err := s.server.Serve()
		s.logger.Error(fmt.Sprintf("stop listening multiplex server: %v", err))
	}()

	<-ctx.Done()
	return fmt.Errorf("server stop listening and serve")
}

func (s *Server) Stop(ctx context.Context) error {
	defer s.server.Close()
	s.grpcServer.GracefulStop()

	return s.httpServer.Shutdown(ctx)
}

//func (s *Server) Handle(w http.ResponseWriter, r *http.Request) {
//	start := time.Now()
//
//	err := s.app.CreateEvent("test test", time.Time{}, "3 days 04:05:06", 2)
//	s.logger.Info(fmt.Sprintf("insert events error: %v", err))
//
//	eventsResult, err := s.app.GetEvents(2)
//	s.logger.Info(fmt.Sprintf("get events: %v and error: %v", eventsResult, err))
//
//	err = s.app.UpdateEvent(2, "test2222")
//	s.logger.Info(fmt.Sprintf("update events error: %v", err))
//
//	eventsResult, err = s.app.GetEvents(2)
//	s.logger.Info(fmt.Sprintf("get events: %v and error: %v", eventsResult, err))
//
//	err = s.app.DeleteEvent(2)
//	s.logger.Info(fmt.Sprintf("delete events error: %v", err))
//
//	msgID := ""
//	w.Header().Add("msgId", msgID)
//	w.Write([]byte("Hello, world"))
//
//	message := fmt.Sprintf(
//		"%s %v %s %s %v %s",
//		r.RemoteAddr, start, r.Method, r.URL.Path, time.Since(start), r.UserAgent())
//	s.logger.Warn(message)
//}
