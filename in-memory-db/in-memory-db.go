package inmemorydb

type inMemory interface {
	SetItem(key string, val string)
	GetItem(key string) string
}

type InMemoryDatabase struct {
	db map[string]string
}

func (i InMemoryDatabase) SetItem(key string, val string) {
	i.db[key] = val
}
func (i InMemoryDatabase) GetItem(key string) string {
	return i.db[key]
}

func New() inMemory {
	database := InMemoryDatabase{}
	database.db = make(map[string]string)
	return database
}
