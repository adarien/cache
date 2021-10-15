package cache

type MyMap map[string]interface{}

type MapStruct struct {
	myMap MyMap
}

func initMap() *MapStruct {
	var m MapStruct
	m.myMap = make(map[string]interface{})
	return &m
}

func New() MapStruct {
	return *initMap()
}

func (m *MapStruct) Set(key string, value interface{}) {
	m.myMap[key] = value
}

func (m *MapStruct) Delete(key string) {
	delete(m.myMap, key)
}

func (m MapStruct) Get(key string) interface{} {
	return m.myMap[key]
}
