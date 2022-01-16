package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	Cache // Remove me after realization.

	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (cache *lruCache) Set(key Key, value interface{}) bool {
	if listItemRefer, ok := cache.items[key]; ok {
		cache.queue.MoveToFront(listItemRefer)
		listItemRefer.Value = cacheItem{key: key, value: value}

		return true
	}

	if cache.queue.Len() >= cache.capacity {
		lastListItem := cache.queue.Back()
		lastCacheItem := lastListItem.Value.(cacheItem)
		delete(cache.items, lastCacheItem.key)
		cache.queue.Remove(lastListItem)
	}

	listItemRefer := cache.queue.PushFront(cacheItem{key: key, value: value})
	cache.items[key] = listItemRefer

	return false
}

func (cache *lruCache) Get(key Key) (interface{}, bool) {
	if listItemRefer, ok := cache.items[key]; ok {
		cache.queue.MoveToFront(listItemRefer)

		return listItemRefer.Value.(cacheItem).value, true
	}

	return nil, false
}

func (cache *lruCache) Clear() {
	cache.queue = NewList()
	cache.items = make(map[Key]*ListItem, cache.capacity)
}
