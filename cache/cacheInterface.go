package cache

type CacheInterface interface {
	Get(key string) (string, error)
	Delete(key string) error
	Set(key string, val string, second int64) error
	Close()
}
