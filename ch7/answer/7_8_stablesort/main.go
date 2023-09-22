package main

import (
	"fmt"
	"reflect"
	"sort"
	"time"
)

type SortableTable struct {
	data     []Row
	sortKeys []string
}

type Row struct {
	// 行的数据
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

func (st SortableTable) Len() int {
	return len(st.data)
}

func (st SortableTable) Less(i, j int) bool {
	for _, key := range st.sortKeys {
		// 根据 sortKeys 中的索引依次比较列头
		// 如果两个行的列头相等，则继续比较下一个列头
		// 如果两个行的列头不相等，则返回比较结果
		fieldI := reflect.ValueOf(st.data[i]).FieldByName(key) //通过反射获取类型字段值
		fieldJ := reflect.ValueOf(st.data[j]).FieldByName(key)
		switch fieldI.Kind() {
		case reflect.String:
			if fieldI.String() != fieldJ.String() {
				return fieldI.String() < fieldJ.String()
			}
		case reflect.Int:
			if fieldI.Int() != fieldJ.Int() {
				return fieldI.Int() < fieldJ.Int()
			}
		case reflect.Struct:
			if fieldI.Interface().(time.Duration) != fieldJ.Interface().(time.Duration) {
				return fieldI.Interface().(time.Duration) < fieldJ.Interface().(time.Duration)
			}
		}
	}
	// 如果所有列头都相等，则返回 false
	return false
}

func (st SortableTable) Swap(i, j int) {
	st.data[i], st.data[j] = st.data[j], st.data[i]
}

func main() {
	st := SortableTable{
		data: []Row{
			{Title: "Song A", Artist: "Artist A", Album: "Album A", Year: 2021, Length: 3 * time.Minute},
			{Title: "Song B", Artist: "Artist B", Album: "Album B", Year: 2020, Length: 4 * time.Minute},
			{Title: "Song C", Artist: "Artist C", Album: "Album C", Year: 2022, Length: 2 * time.Minute},
		},
		sortKeys: []string{"Year", "Title", "Artist"},
	}

	// 使用 SortableTable 进行排序
	fmt.Println(st.data)
	sort.Sort(st)
	fmt.Println(st.data)
}
