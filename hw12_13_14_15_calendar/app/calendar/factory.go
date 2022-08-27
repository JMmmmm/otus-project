package calendar

import (
	"context"
	"fmt"
	"github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/pkg/logger"

	domain "github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/domain/calendarevent"
	memoryrepository "github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/internal/repository/calendarevent/memory"
	sqlrepository "github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/internal/repository/calendarevent/sql"
)

const (
	dbTypeSQL    = "sql"
	dbTypeMemory = "memory"
)

func CreateApp(ctx context.Context, config Config, logger logger.Logger) (*App, error) {
	var err error
	var calendarEventRepository domain.CalendarEventRepository

	switch config.DB.DBType {
	case dbTypeSQL:
		calendarEventRepository, err = sqlrepository.NewCalendarEventRepository(ctx, logger, config.PSQL.DSN)
		if err != nil {
			return nil, fmt.Errorf("can not create sql connection: %w", err)
		}
	case dbTypeMemory:
		calendarEventRepository = memoryrepository.NewCalendarEventRepository()
	default:
		return nil, fmt.Errorf("undefined db type: %s", config.DB.DBType)
	}

	return New(logger, calendarEventRepository), nil
}
