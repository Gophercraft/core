package wdb

import "reflect"

const lrusize = 64

type lruanswer struct {
	indexType BucketKeyType
	key       any
	// A pointer to an underlying data type
	pointer reflect.Value
	// closer to head.
	prev *lruanswer
	// closer to tail
	next *lruanswer
}

func (b *Bucket) cacheAnswer(indexType BucketKeyType, key any, storage reflect.Value) {
	first := b.lruHead == nil

	previousHead := b.lruHead

	b.lruHead = &lruanswer{
		indexType: indexType,
		key:       key,
		pointer:   storage.Elem(),
		prev:      nil,
	}

	if previousHead != nil {
		previousHead.prev = b.lruHead
		b.lruHead.next = previousHead
	}

	if first {
		b.lruTail = b.lruHead
	}

	b.lrusize++

	if b.lrusize == lrusize {
		b.lruTail = b.lruTail.prev
		b.lrusize--
	}
}

func (b *Bucket) findInMemoryAnswer(indexType BucketKeyType, key any, storage reflect.Value) bool {
	if b.lruHead == nil {
		return false
	}

	ans := b.lruHead

	for ans != nil {
		if ans.indexType == indexType && ans.key == key {
			// remove ans from linked list.
			beforeAns := ans.prev
			afterAns := ans.next
			ans.prev = nil
			ans.next = nil

			if beforeAns != nil {
				beforeAns.next = afterAns
			}

			if afterAns != nil {
				afterAns.prev = beforeAns
			}

			// move ans to head of linked list
			previousLruHead := b.lruHead
			b.lruHead = ans
			b.lruHead.next = previousLruHead

			storage.Elem().Set(ans.pointer)
			return true
		}
		ans = ans.next
	}

	return false
}
