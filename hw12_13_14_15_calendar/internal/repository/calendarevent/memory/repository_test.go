package memoryrepository

import (
	"testing"

	domain "github.com/JMmmmm/otus-project/hw12_13_14_15_calendar/domain/calendarevent"
	"github.com/stretchr/testify/require"
)

func TestRepository(t *testing.T) {
	t.Run("init", func(t *testing.T) {
		repository := NewCalendarEventRepository()
		test := domain.CalendarEventEntity{ID: "123-41234-1245", Title: "Test title", UserID: 2}
		test2 := domain.CalendarEventEntity{ID: "123-41234-7777", Title: "Test title 2", UserID: 1}
		test3 := domain.CalendarEventEntity{ID: "123-41234-8888", Title: "Test title 2", UserID: 3}
		err := repository.InsertEntities([]domain.CalendarEventEntity{test, test2, test3})
		require.NoError(t, err)

		test3 = domain.CalendarEventEntity{ID: "123-41234-8888", Title: "Test title 3", UserID: 3}
		err = repository.Update(test3)

		var val []domain.CalendarEventEntity
		val, err = repository.GetEvents(test.UserID)
		require.NoError(t, err)
		require.Equal(t, []domain.CalendarEventEntity{test}, val)

		val, err = repository.GetEvents(test2.UserID)
		require.NoError(t, err)
		require.Equal(t, []domain.CalendarEventEntity{test2}, val)

		val, err = repository.GetEvents(test3.UserID)
		require.NoError(t, err)
		require.Equal(t, []domain.CalendarEventEntity{test3}, val)

		err = repository.Delete(test2.ID)
		require.NoError(t, err)

		val, err = repository.GetEvents(test.UserID)
		require.NoError(t, err)
		require.Equal(t, []domain.CalendarEventEntity{test}, val)

		val, err = repository.GetEvents(test2.UserID)
		require.Error(t, err)
		require.Equal(t, []domain.CalendarEventEntity{}, val)

		val, err = repository.GetEvents(test3.UserID)
		require.NoError(t, err)
		require.Equal(t, []domain.CalendarEventEntity{test3}, val)
	})
}
