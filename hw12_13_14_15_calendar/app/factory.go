package app

import (
	"context"
	"fmt"
	"github.com/fixme_my_friend/hw12_13_14_15_calendar/domain"
	"github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/logger"
	memoryrepository "github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/repository/calendarevent/memory"
	sqlrepository "github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/repository/calendarevent/sql"
	memorystorage "github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/storage/sql"
)

const (
	dbTypeSql    = "sql"
	dbTypeMemory = "memory"
)

func CreateApp(ctx context.Context, config Config, logger logger.Logger) (*App, error) {
	var err error
	var calendarEventRepository domain.CalendarEventRepository

	switch config.DB.DbType {
	case dbTypeSql:
		storage := sqlstorage.New()
		err = storage.Connect(ctx, config.PSQL.DSN)
		if err != nil {
			return nil, fmt.Errorf("can not create sql connection: %v", err)
		}
		calendarEventRepository = sqlrepository.NewCalendarEventRepository(storage, logger)
	case dbTypeMemory:
		storage := memorystorage.New()
		calendarEventRepository = memoryrepository.NewCalendarEventRepository(storage, logger)
	default:
		return nil, fmt.Errorf("undefined db type: %s", config.DB.DbType)
	}

	return New(logger, calendarEventRepository), nil
}
