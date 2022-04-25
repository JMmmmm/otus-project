package domain

type CalendarEventRepository interface {
	GetEvents(userID int) ([]CalendarEventEntity, error)
	InsertEntities(entities []CalendarEventEntity) error
	Update(entity CalendarEventEntity) error
	Delete(id string) error
}
