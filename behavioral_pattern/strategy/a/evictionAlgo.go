package a

type EvictionAlgo interface {
	evict(c *Cache)
}
