package domain

type CalendarEventRepository interface {
	GetEvents(userId int) ([]CalendarEventEntity, error)
	Insert(entities []CalendarEventEntity) error
	Update(entity CalendarEventEntity) error
	Delete(userId int) error
}
