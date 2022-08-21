package api

import (
	"context"
	"github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/app/calendar"
	"log"
	"testing"

	"github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/internal/logger"
	calendar_event_api "github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/pkg/calendar-event"
	"github.com/stretchr/testify/require"
)

func getController() *Controller {
	config, err := calendar.NewConfig("../configs/calendar_config.toml")
	if err != nil {
		log.Fatalf("Can not create config, %v", err)
	}
	logg, err := logger.NewAppLogger(config.Logger.Level, config.Logger.OutputPath)
	if err != nil {
		log.Fatalf("Can not create logger: %v", err)
	}

	ctx := context.Background()

	calendar, err := calendar.CreateApp(ctx, config, logg)
	if err != nil {
		log.Fatalf("Can not create App: %v", err)
	}

	return &Controller{
		App: calendar,
	}
}

func TestController_GetEvents(t *testing.T) {
	controller := getController()

	testCases := []struct {
		name        string
		req         *calendar_event_api.GetUserEventsRequest
		response    *calendar_event_api.Events
		expectedErr bool
		message     string
	}{
		{
			name: "req ok",
			req:  &calendar_event_api.GetUserEventsRequest{UserId: 1},
			response: &calendar_event_api.Events{
				Content: []*calendar_event_api.Event{
					{
						Id:     "1",
						UserId: 1,
						Title:  "test",
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		testCase := tc
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			response, err := controller.GetEvents(ctx, testCase.req)
			if err != nil {
				log.Fatalf("response error: %v", err)
			}
			if testCase.expectedErr {
				require.Error(t, err, testCase.message)
			} else {
				require.Equal(t, response.Content, testCase.response.Content)
			}
		},
		)
	}
}

func TestCRUD(t *testing.T) {
	controller := getController()

	_, err := controller.CreateEvent(context.Background(), &calendar_event_api.CreateEventRequest{
		UserId:   5,
		Title:    "test",
		Duration: "3 days 04:05:06",
	})
	require.NoError(t, err)

	event, err := controller.GetEvents(context.Background(), &calendar_event_api.GetUserEventsRequest{UserId: 5})
	require.NoError(t, err)

	_, err = controller.UpdateEvent(context.Background(), &calendar_event_api.UpdateEventRequest{
		Id:    event.Content[0].Id,
		Title: "new test",
	})
	require.NoError(t, err)

	_, err = controller.DeleteEvent(context.Background(), &calendar_event_api.DeleteEventRequest{Id: event.Content[0].Id})
	require.NoError(t, err)
}
