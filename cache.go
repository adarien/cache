package cache

type Cache struct {
	currentCache map[string]interface{}
}

func New() *Cache {
	return &Cache{
		currentCache: make(map[string]interface{}),
	}
}

func (m *Cache) Set(key string, value interface{}) {
	m.currentCache[key] = value
}

func (m *Cache) Delete(key string) {
	delete(m.currentCache, key)
}

func (m Cache) Get(key string) interface{} {
	value, exists := m.currentCache[key]
	if exists {
		return value
	}
	return "key not found"
}
