package domain

type CalendarEventRepository interface {
	GetEvents(userID int) ([]CalendarEventEntity, error)
	Insert(entities []CalendarEventEntity) error
	Update(entity CalendarEventEntity) error
	Delete(userID int) error
}
