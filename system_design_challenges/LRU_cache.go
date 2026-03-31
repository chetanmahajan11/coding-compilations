package systemdesignchallenges

type LRUCache struct {
	capacity int
	cache    map[int]int
	order    []int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]int),
		order:    []int{},
	}
}

// Get() returns the value of the key if the key exists, otherwise return -1.
func (this *LRUCache) Get(key int) int {
	value, present := this.cache[key]
	if !present {
		return -1
	}

	// move to recent=> remove from 'order' and append to last(last=most recent)
	this.reorder(key)
	return value
}

// Put() updates the value of the key if the key exists.
// Otherwise, add the key-value pair to the cache. If the number of keys exceeds the capacity from this operation, evict the least recently used key.
func (this *LRUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}
	if _, present := this.cache[key]; present {

		// stores key-value
		this.cache[key] = value
		// reorder
		this.reorder(key)
		return
	}

	// capacity exceeds, evict  least recently used key.
	if this.capacity == len(this.cache) {
		leastRecentlyUsed := this.order[0]
		this.order = this.order[1:]
		delete(this.cache, leastRecentlyUsed)
	}

	// stores key-value
	this.cache[key] = value
	this.order = append(this.order, key)
}

// reorder shifts key to last
func (this *LRUCache) reorder(key int) {
	for index, value := range this.order {
		if value == key {
			this.order = append(this.order[:index], this.order[index+1:]...)
			this.order = append(this.order, key)
			return
		}
	}
}

/*
 most recent = end of slice
 least recent = start of slice
 on Get: move key to end
 on Push: evict first, if full
*/
