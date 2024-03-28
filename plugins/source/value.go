package source

import "sync"

type valueCache struct {
	updated int64
	data    map[string]*item
	lock    sync.RWMutex
}

type item struct {
	value   string
	updated int64
}

func newValueCache(updated int64) *valueCache {
	return &valueCache{
		updated: updated,
		data:    make(map[string]*item),
	}
}

func (vc *valueCache) getUpdated() int64 {
	vc.lock.RLock()
	defer vc.lock.RUnlock()

	return vc.updated
}

func (vc *valueCache) setUpdated(updated int64) {
	vc.lock.Lock()
	defer vc.lock.Unlock()

	vc.updated = updated
}

func (vc *valueCache) getTagValue(id string) (*item, bool) {
	vc.lock.RLock()
	defer vc.lock.RUnlock()

	if v, ok := vc.data[id]; ok {
		return v, true
	}

	return nil, false
}

func (vc *valueCache) setTagValue(id string, value string, updated int64) (sync bool) {
	vc.lock.Lock()
	defer vc.lock.Unlock()

	if v, ok := vc.data[id]; ok {
		if v.value != value {
			v.value = value
			v.updated = updated
			sync = true
		}

		return
	}

	vc.data[id] = &item{value: value, updated: updated}

	return true
}
