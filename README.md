# golang-lib
Library which realize cache pattern
There are next methods
- Set(key string, value interface{})  write data to cache
- Get(key string) return saved data. Nil if not exist
- Delete(key string) remove data from cache
- New() create new cache

Example

func main() {
	cache := cache.New()

	cache.Set("field1", 43)
	cache.Set("field2", "test string")

	fmt.Println(cache.Get("field1"))
	fmt.Println(cache.Get("field2"))

	cache.Delete("field2")
	fmt.Println(cache.Get("field2"))

}
