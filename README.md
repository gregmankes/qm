# QueueMap

A map that maintains insertion ordering.

## Fast QueueMap
### Performance
* Add
  * O(1)
* Fetch
  * O(1)
* Remove
  * O(1)
* Dump
  * O(n)

### Space Characteristics
Need to store the key in the value structure to use on dump. Space is O(1.5n)

## Simple QueueMap
### Performance
* Add
  * O(1)
* Fetch
  * O(1)
* Remove
  * O(1)
* Dump
  * O(2n + nlogn) => O(nlogn)

### Space Characteristics
O(n)

## Side/Speed Tradeoff
From the two implementations we can see that dump method is faster with the fast QueueMap, but uses slightly more space.

The fast QueueMap creates a doubly linked list queue between all of the "value" nodes inserted in the map. It allows an O(n) dump of all the nodes, but requires the space of having the key stored in the node.

The simple QueueMap takes advantage of just keeping count of when the key was inserted, but you need to pull all keys out and then sort to get the insertion order.

## Usage

```
// Two types:
// FastType
// SimpleType

queuemap.New(queuemap.FastType)
```