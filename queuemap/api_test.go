package queuemap_test

import (
	"reflect"
	"testing"

	"github.com/gregmankes/qm/queuemap"
)

const (
	nonExistentKey = "foobarbaz"
)

func getQueueMaps() map[string]queuemap.QueueMap {
	// instantiate a fast and simple queuemap
	simpleQueueMap, _ := queuemap.New(queuemap.SimpleType)
	fastQueueMap, _ := queuemap.New(queuemap.FastType)

	// add them to a map for table driven testing
	return map[string]queuemap.QueueMap{
		"simple": simpleQueueMap,
		"fast":   fastQueueMap,
	}
}

func TestInsertionOrder(t *testing.T) {
	testcases := [][]queuemap.KeyValue{
		[]queuemap.KeyValue{
			queuemap.KeyValue{
				Key:   "foo1",
				Value: "bar1",
			},
			queuemap.KeyValue{
				Key:   "foo2",
				Value: "bar2",
			},
			queuemap.KeyValue{
				Key:   "foo3",
				Value: "bar3",
			},
			queuemap.KeyValue{
				Key:   "foo4",
				Value: "bar4",
			},
			queuemap.KeyValue{
				Key:   "foo5",
				Value: "bar5",
			},
		},
		[]queuemap.KeyValue{
			queuemap.KeyValue{
				Key:   "bar1",
				Value: "foo1",
			},
			queuemap.KeyValue{
				Key:   "bar2",
				Value: "foo2",
			},
			queuemap.KeyValue{
				Key:   "bar3",
				Value: "foo3",
			},
			queuemap.KeyValue{
				Key:   "bar4",
				Value: "foo4",
			},
			queuemap.KeyValue{
				Key:   "bar5",
				Value: "foo5",
			},
		},
	}

	// get map of queuemaps
	queuemaps := getQueueMaps()

	for name, qm := range queuemaps {
		for i, tc := range testcases {
			for _, kv := range tc {
				qm.Add(kv.Key, kv.Value)
			}
			if !reflect.DeepEqual(qm.Dump(), tc) {
				t.Errorf("Testcase %d failed for %s queuemap. Expected: %v, actual: %v", i, name, tc, qm.Dump())
			}
			for _, kv := range tc {
				qm.Remove(kv.Key)
			}
		}
	}
}

func TestRemovalInsertionOrder(t *testing.T) {
	type testcase struct {
		In  []queuemap.KeyValue
		Out []queuemap.KeyValue
	}
	testcases := []testcase{
		testcase{
			In: []queuemap.KeyValue{
				queuemap.KeyValue{
					Key:   "foo1",
					Value: "bar1",
				},
				queuemap.KeyValue{
					Key:   "foo2",
					Value: "bar2",
				},
				queuemap.KeyValue{
					Key:   "foo3",
					Value: "bar3",
				},
				queuemap.KeyValue{
					Key:   "foo4",
					Value: "bar4",
				},
				queuemap.KeyValue{
					Key:   "foo5",
					Value: "bar5",
				},
			},
			Out: []queuemap.KeyValue{
				queuemap.KeyValue{
					Key:   "foo1",
					Value: "bar1",
				},
				queuemap.KeyValue{
					Key:   "foo2",
					Value: "bar2",
				},
				queuemap.KeyValue{
					Key:   "foo4",
					Value: "bar4",
				},
				queuemap.KeyValue{
					Key:   "foo5",
					Value: "bar5",
				},
			},
		},
		testcase{
			In: []queuemap.KeyValue{
				queuemap.KeyValue{
					Key:   "bar1",
					Value: "foo1",
				},
				queuemap.KeyValue{
					Key:   "bar2",
					Value: "foo2",
				},
				queuemap.KeyValue{
					Key:   "bar3",
					Value: "foo3",
				},
				queuemap.KeyValue{
					Key:   "bar4",
					Value: "foo4",
				},
				queuemap.KeyValue{
					Key:   "bar5",
					Value: "foo5",
				},
			},
			Out: []queuemap.KeyValue{
				queuemap.KeyValue{
					Key:   "bar1",
					Value: "foo1",
				},
				queuemap.KeyValue{
					Key:   "bar2",
					Value: "foo2",
				},
				queuemap.KeyValue{
					Key:   "bar4",
					Value: "foo4",
				},
				queuemap.KeyValue{
					Key:   "bar5",
					Value: "foo5",
				},
			},
		},
	}

	// get map of queuemaps
	queuemaps := getQueueMaps()

	for name, qm := range queuemaps {
		for i, tc := range testcases {
			for _, kv := range tc.In {
				qm.Add(kv.Key, kv.Value)
			}
			val := qm.Remove(nonExistentKey)
			if val != "" {
				t.Error("removed value incorrect")
			}
			val = qm.Remove(tc.In[2].Key)
			if val != tc.In[2].Value {
				t.Error("removed value incorrect")
			}
			qm.Remove(tc.In[2].Key)
			if !reflect.DeepEqual(qm.Dump(), tc.Out) {
				t.Errorf("Testcase %d failed for queuemap %s. Expected: %v, actual: %v", i, name, tc.Out, qm.Dump())
			}
			for _, kv := range tc.In {
				qm.Remove(kv.Key)
			}
		}
	}
}

func TestDupKey(t *testing.T) {
	type testcase struct {
		In  []queuemap.KeyValue
		Out []queuemap.KeyValue
	}
	testcases := []testcase{
		testcase{
			In: []queuemap.KeyValue{
				queuemap.KeyValue{
					Key:   "foo1",
					Value: "bar1",
				},
				queuemap.KeyValue{
					Key:   "foo2",
					Value: "bar2",
				},
				queuemap.KeyValue{
					Key:   "foo1",
					Value: "baz1",
				},
			},
			Out: []queuemap.KeyValue{
				queuemap.KeyValue{
					Key:   "foo2",
					Value: "bar2",
				},
				queuemap.KeyValue{
					Key:   "foo1",
					Value: "baz1",
				},
			},
		},
	}

	// get map of queuemaps
	queuemaps := getQueueMaps()

	for name, qm := range queuemaps {
		for i, tc := range testcases {
			for _, kv := range tc.In {
				qm.Add(kv.Key, kv.Value)
			}
			if !reflect.DeepEqual(qm.Dump(), tc.Out) {
				t.Errorf("Testcase %d for %s queuemap, got %v expected %v", i, name, qm.Dump(), tc.Out)
			}
			for _, kv := range tc.In {
				qm.Remove(kv.Key)
			}
		}
	}
}

func TestFetch(t *testing.T) {
	type testcase struct {
		In  []queuemap.KeyValue
		Key string
		Out string
	}
	testcases := []testcase{
		testcase{
			In: []queuemap.KeyValue{
				queuemap.KeyValue{
					Key:   "foo1",
					Value: "bar1",
				},
				queuemap.KeyValue{
					Key:   "foo2",
					Value: "bar2",
				},
			},
			Key: "foo1",
			Out: "bar1",
		},
		testcase{
			In: []queuemap.KeyValue{
				queuemap.KeyValue{
					Key:   "foo1",
					Value: "bar1",
				},
				queuemap.KeyValue{
					Key:   "foo2",
					Value: "bar2",
				},
			},
			Key: "foo3",
			Out: "",
		},
	}

	// get map of queuemaps
	queuemaps := getQueueMaps()

	for name, qm := range queuemaps {
		for i, tc := range testcases {
			for _, kv := range tc.In {
				qm.Add(kv.Key, kv.Value)
			}
			if got := qm.Fetch(tc.Key); got != tc.Out {
				t.Errorf("Testcase %d failed for %s queuemap, got %v expected %v", i, name, got, tc.Out)
			}
			for _, kv := range tc.In {
				qm.Remove(kv.Key)
			}
		}
	}
}
