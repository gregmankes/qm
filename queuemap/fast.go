package queuemap

type fastQueueMap struct {
	storage map[string]*queueNode
	queue   *queue
}

func newFastQueueMap() *fastQueueMap {
	return &fastQueueMap{
		storage: make(map[string]*queueNode),
		queue:   newQueue(),
	}
}

func (s *fastQueueMap) Add(key, val string) {
	qn := &queueNode{
		KeyValue: KeyValue{
			Key:   key,
			Value: val,
		},
	}
	if _, ok := s.storage[key]; ok {
		s.Remove(key)
	}
	s.storage[key] = qn
	s.queue.Push(qn)
}

func (s *fastQueueMap) Fetch(key string) string {
	if val, ok := s.storage[key]; ok && val != nil {
		return val.KeyValue.Value
	}
	return ""
}

func (s *fastQueueMap) Remove(key string) string {
	qn := s.storage[key]
	if qn == nil {
		return ""
	}
	s.queue.Remove(qn)
	delete(s.storage, key)
	return qn.KeyValue.Value
}

func (s *fastQueueMap) Dump() []KeyValue {
	cur := s.queue.Head
	kvs := []KeyValue{}
	for cur != nil {
		kvs = append(kvs, cur.KeyValue)
		cur = cur.Next
	}
	return kvs
}
