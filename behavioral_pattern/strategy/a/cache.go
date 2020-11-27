package a


type Cache struct {
	storage map[string]string
	evicttionAlgo EvictionAlgo
	capacity int
	maxCapacity int
}

/*InitCache is function*/
func InitCache(e EvictionAlgo) *Cache{
	storage :=make(map[string]string)
	return &Cache{
		storage:storage,
		evicttionAlgo: e,
		capacity: 0,
		maxCapacity:2,
	}
}

func (c *Cache) SetEvictionAlgo(e EvictionAlgo){
	c.evicttionAlgo=e
}

func (c *Cache)Add(key, value string){
	if c.capacity==c.maxCapacity{
		c.evicttionAlgo.evict(c)
		c.capacity--
	}
	c.capacity++
	c.storage[key]=value
}