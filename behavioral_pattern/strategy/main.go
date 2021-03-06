package main


import s "behavior/strategy/a"

func main(){
	lfu := &s.LFU{}
	cache := s.InitCache(lfu)
	cache.Add("a","1")
	cache.Add("b","2")
	cache.Add("c","3")

	lru := &s.LRU{}
	cache.SetEvictionAlgo(lru)
	cache.Add("d","4")

	fifo :=&s.FIFO{}
	cache.SetEvictionAlgo(fifo)
	cache.Add("e", "5")
}
