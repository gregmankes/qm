package queuemap

import (
	"fmt"
)

// Type the type for switching on creating the queuemap
type Type string

const (
	// SimpleType the key to use to create a simple QueueMap
	SimpleType Type = "SIMPLE"

	// FastType the key to use to create a fast QueueMap
	FastType Type = "FAST"
)

// QueueMap a map that preserves the insertion order of the keys
type QueueMap interface {
	Add(string, string)
	Fetch(string) string
	Remove(string) string
	Dump() []KeyValue
}

// New Constructor, use the constants to initialize a QueueMap
func New(mapType Type) (QueueMap, error) {
	switch mapType {
	case SimpleType:
		return newSimpleQueueMap(), nil
	case FastType:
		return newFastQueueMap(), nil
	default:
		return nil, fmt.Errorf("Error creating the QueueMap: %s is not a correct map type", mapType)
	}
}

// KeyValue the return value of the dump method
type KeyValue struct {
	Key   string
	Value string
}
