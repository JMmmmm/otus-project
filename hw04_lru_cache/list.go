package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	firstItem *ListItem
	lastItem  *ListItem
	len       int
}

func (list *list) PushFront(v interface{}) *ListItem {
	ListItemStruct := ListItem{Value: v}
	if list.firstItem == nil {
		list.firstItem = &ListItemStruct
		list.lastItem = &ListItemStruct
	} else {
		list.insertBefore(list.firstItem, &ListItemStruct)
	}

	list.len++

	return &ListItemStruct
}

func (list *list) MoveToFront(i *ListItem) {
	if list.firstItem == i {
		return
	}
	if i == list.lastItem {
		list.lastItem = i.Prev
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	ListItemStruct := ListItem{Value: i.Value}
	list.insertBefore(list.firstItem, &ListItemStruct)
}

func (list *list) PushBack(v interface{}) *ListItem {
	if list.lastItem == nil {
		return list.PushFront(v)
	}
	ListItemStruct := ListItem{Value: v}
	list.insertAfter(list.lastItem, &ListItemStruct)

	list.len++

	return &ListItemStruct
}

func (list *list) Remove(i *ListItem) {
	if i.Prev == nil {
		list.firstItem = i.Next
	} else {
		i.Prev.Next = i.Next
	}
	if i.Next == nil {
		list.lastItem = i.Prev
	} else {
		i.Next.Prev = i.Prev
	}

	list.len--
}

func (list *list) Len() int {
	return list.len
}

func (list *list) Front() *ListItem {
	return list.firstItem
}

func (list *list) Back() *ListItem {
	return list.lastItem
}

func NewList() List {
	return new(list)
}

func (list *list) insertAfter(item *ListItem, newItem *ListItem) {
	newItem.Prev = item
	if item.Next == nil {
		list.lastItem = newItem
	} else {
		newItem.Next = item.Next
		item.Next.Prev = newItem
	}

	item.Next = newItem
}

func (list *list) insertBefore(item *ListItem, newItem *ListItem) {
	newItem.Next = item
	if item.Prev == nil {
		list.firstItem = newItem
	} else {
		newItem.Prev = item.Prev
		item.Prev.Next = newItem
	}
	item.Prev = newItem
}
