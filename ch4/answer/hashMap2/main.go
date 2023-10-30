/*
模拟map的底层实现-hashTable
修复后的代码，使用开放定址法解决哈希冲突，并且支持动态调整哈希表大小
*/
package main

import (
	"fmt"
	"hash/fnv"
)

const (
	InitialSize = 1
	LoadFactor  = 0.75
)

type Entry struct {
	Key   string
	Value int
}

type HashTable struct {
	Table    []*Entry //存的是地址切片
	Used     uint32
	Capacity uint32
}

func NewHashTable() *HashTable {
	return &HashTable{
		Table:    make([]*Entry, InitialSize),
		Used:     0,
		Capacity: InitialSize,
	}
}

func (ht *HashTable) Insert(key string, value int) {
	if float64(ht.Used)/float64(ht.Capacity) >= LoadFactor {
		ht.resize()
	}

	index := ht.hash(key)

	for ht.Table[index] != nil {
		if ht.Table[index].Key == key {
			ht.Table[index].Value = value
			return
		}
		index = (index + 1) % ht.Capacity
	}

	ht.Table[index] = &Entry{Key: key, Value: value}
	ht.Used++
}

func (ht *HashTable) Get(key string) (int, bool) {
	index := ht.hash(key)

	for ht.Table[index] != nil {
		if ht.Table[index].Key == key {
			return ht.Table[index].Value, true
		}
		index = (index + 1) % ht.Capacity
	}

	return 0, false
}

func (ht *HashTable) hash(key string) uint32 {
	hash := fnv.New32a()
	hash.Write([]byte(key))
	return hash.Sum32() % uint32(ht.Capacity)
}

func (ht *HashTable) resize() {
	newCapacity := ht.Capacity * 2
	newTable := make([]*Entry, newCapacity)

	for _, entry := range ht.Table {
		if entry != nil {
			index := ht.hash(entry.Key)

			for newTable[index] != nil {
				index = (index + 1) % newCapacity
			}

			newTable[index] = entry
		}
	}

	ht.Table = newTable
	ht.Capacity = newCapacity
}

func main() {
	ht := NewHashTable()

	ht.Insert("apple", 5)
	ht.Insert("banana", 10)
	ht.Insert("orange", 15)

	value, found := ht.Get("banana")
	if found {
		fmt.Println("Value:", value)
	} else {
		fmt.Println("Key not found")
	}
}
