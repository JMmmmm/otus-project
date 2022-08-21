package internalhttp

import (
	"context"
	"fmt"
	"github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/app/calendar"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/api"
	"github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/internal/logger"
	calendar_event_api "github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/pkg/calendar-event"
	"github.com/felixge/httpsnoop"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	logger       logger.Logger
	server       cmux.CMux
	httpServer   *http.Server
	grpcServer   *grpc.Server
	httpListener net.Listener
	grpcListener net.Listener
}

func withLogger(handler http.Handler, logger logger.Logger) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		m := httpsnoop.CaptureMetrics(handler, writer, request)
		logger.Info(fmt.Sprintf(
			"%s-- %v -- http[%d]-- %s -- %s -- %s\n",
			request.RemoteAddr,
			time.Now(),
			m.Code,
			m.Duration,
			request.URL.Path,
			request.UserAgent(),
		))
	})
}

func NewServer(logger logger.Logger, app calendar.Application, addr string) *Server {
	grpcSever := grpc.NewServer()

	calendar_event_api.RegisterCalendarEventServer(grpcSever, &api.Controller{
		App: app,
	})
	mux := runtime.NewServeMux()
	err := calendar_event_api.RegisterCalendarEventHandlerFromEndpoint(
		context.Background(),
		mux,
		addr,
		[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
	)
	if err != nil {
		log.Fatal(err)
	}
	server := &http.Server{
		Handler: withLogger(mux, logger),
	}
	l, err := net.Listen("tcp", addr)
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
