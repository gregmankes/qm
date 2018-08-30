package queuemap

import (
	"sort"
)

type value struct {
	val   string
	order int
}

type simpleQueueMap struct {
	storage map[string]value
	count   int
}

func newSimpleQueueMap() *simpleQueueMap {
	return &simpleQueueMap{
		storage: make(map[string]value),
	}
}

func (s *simpleQueueMap) Add(key, val string) {
	if _, ok := s.storage[key]; ok {
		s.Remove(key)
	}
	s.storage[key] = value{
		val:   val,
		order: s.count,
	}
	s.count++
}

func (s *simpleQueueMap) Fetch(key string) string {
	return s.storage[key].val
}

func (s *simpleQueueMap) Remove(key string) string {
	val := s.storage[key].val
	delete(s.storage, key)
	return val
}

func (s *simpleQueueMap) Dump() []KeyValue {
	type wrapper struct {
		KV    KeyValue
		order int
	}
	wrappers := []wrapper{}
	for key, val := range s.storage {
		wrappers = append(wrappers, wrapper{
			KV: KeyValue{
				Key:   key,
				Value: val.val,
			},
			order: val.order,
		})
	}
	sort.Slice(wrappers, func(i, j int) bool {
		return wrappers[i].order < wrappers[j].order
	})
	kvs := []KeyValue{}
	for _, w := range wrappers {
		kvs = append(kvs, w.KV)
	}
	return kvs
}
