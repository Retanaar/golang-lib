package cache

type Cache struct {
	buf map[string]interface{}
}

func (c Cache) Set(key string, value interface{}) {
	c.buf[key] = value
}

func (c Cache) Get(key string) interface{} {
	return c.buf[key]
}

func (c Cache) Delete(key string) {
	delete(c.buf, key)
}

func New() Cache {
	return Cache{buf: make(map[string]interface{})}
}
