package memorystorage_test

import (
	domain "github.com/fixme_my_friend/hw12_13_14_15_calendar/domain/calendarevent"
	memorystorage "github.com/fixme_my_friend/hw12_13_14_15_calendar/internal/storage/memory"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStorage(t *testing.T) {
	t.Run("init", func(t *testing.T) {
		s := memorystorage.New()
		test := &domain.CalendarEventEntity{ID: "123-41234-1245", Title: "Test title"}
		test2 := &domain.CalendarEventEntity{ID: "123-41234-7777", Title: "Test title 2"}
		test3 := &domain.CalendarEventEntity{ID: "123-41234-8888", Title: "Test title 3"}
		s.Insert("test", test)
		s.Insert("test2", test2)
		s.Update("test3", test3)
		var (
			val domain.CalendarEventEntity
			err error
		)
		val, err = s.Find("test")
		require.Equal(t, nil, err)
		require.Equal(t, *test, val)
		val, err = s.Find("test2")
		require.Equal(t, nil, err)
		require.Equal(t, *test2, val)
		val, err = s.Find("test3")
		require.Equal(t, nil, err)
		require.Equal(t, *test3, val)
		s.Delete("test2")
		val, err = s.Find("test")
		require.Equal(t, nil, err)
		require.Equal(t, *test, val)
		val, err = s.Find("test2")
		require.NotEqual(t, nil, err)
		require.Equal(t, domain.CalendarEventEntity{}, val)
		val, err = s.Find("test3")
		require.Equal(t, nil, err)
		require.Equal(t, *test3, val)
	})
	t.Run("get list", func(t *testing.T) {
		s := memorystorage.New()
		test := &domain.CalendarEventEntity{ID: "123-41234-1245", Title: "Test title"}
		test2 := &domain.CalendarEventEntity{ID: "123-41234-7777", Title: "Test title 2"}
		s.Insert("test", test)
		s.Insert("test2", test2)
		require.Equal(t, 2, len(s.Get(0, 20)))
		s.Delete("test")
		require.Equal(t, 1, len(s.Get(0, 20)))
	})
}
