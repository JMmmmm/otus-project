package app

import (
	"context"
	"fmt"

	domain "github.com/fixme_my_friend/hw12_13_14_15_calendar/domain/calendarevent"
	"github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/logger"
	memoryrepository "github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/repository/calendarevent/memory"
	sqlrepository "github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/repository/calendarevent/sql"
	memorystorage "github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/storage/sql"
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
		storage := sqlstorage.New()
		err = storage.Connect(ctx, config.PSQL.DSN)
		if err != nil {
			return nil, fmt.Errorf("can not create sql connection: %w", err)
		}
		calendarEventRepository = sqlrepository.NewCalendarEventRepository(storage, logger)
	case dbTypeMemory:
		storage := memorystorage.New()
		calendarEventRepository = memoryrepository.NewCalendarEventRepository(storage)
	default:
		return nil, fmt.Errorf("undefined db type: %s", config.DB.DBType)
	}

	return New(logger, calendarEventRepository), nil
}
