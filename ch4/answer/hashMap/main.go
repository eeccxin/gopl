/*
模拟map的底层实现-hashTable
开放定址法
问题：超过TableSize时会死循环
*/
package main

import "fmt"

const TableSize = 1

type Entry struct {
	Key   string
	Value int
}

type HashTable struct {
	Table [TableSize]*Entry
}

func (ht *HashTable) Insert(key string, value int) {
	index := hash(key)

	for ht.Table[index] != nil {
		index = (index + 1) % TableSize
	}

	ht.Table[index] = &Entry{Key: key, Value: value}
}

func (ht *HashTable) Get(key string) (int, bool) {
	index := hash(key)

	for ht.Table[index] != nil {
		if ht.Table[index].Key == key {
			return ht.Table[index].Value, true
		}
		index = (index + 1) % TableSize
	}

	return 0, false
}

func hash(key string) int {
	sum := 0
	for _, ch := range key {
		sum += int(ch)
	}
	return sum % TableSize
}

func main() {
	ht := HashTable{}

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
