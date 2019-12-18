package slice

import "reflect"

type Int64Slice []int64

func (p Int64Slice) Len() int           { return len(p) }
func (p Int64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// 向一个 []int64 的指定位置插入一个 int64
// s := []int64{1,2,3,4,5}
// s = insertSlice(s,2,6)  // [1 2 6 3 4 5]
func InsertInt64Slice(s []int64, pos int64, value int64) []int64 {
	s = append(s[:pos+1], s[pos:]...)
	s[pos] = value
	return s
}

// Int64SliceIntersect 求2个int64数组的交集, 结果跟s1顺序一致
func Int64SliceIntersect(s1, s2 []int64) []int64 {
	var s2Map = make(map[int64]bool, len(s2))
	for _, v := range s2 {
		s2Map[v] = true
	}

	var (
		i   int
		res = make([]int64, len(s1))
	)
	for _, v1 := range s1 {
		if _, ok := s2Map[v1]; ok {
			res[i] = v1
			i++
		}
	}

	return res[:i]
}

//FindInSlice 查找元素在切片中位置，同mysql中内置函数find_in_set()
func FindInSlice(one interface{}, all interface{}) (pos int) {
	if all == nil || one == nil {
		return
	}

	av := reflect.ValueOf(all)
	if av.Kind() == reflect.Ptr {
		av = av.Elem()
	}

	if av.Kind() != reflect.Slice {
		return
	}

	for i := 0; i < av.Len(); i++ {
		if av.Index(i).Interface() == one {
			pos = i + 1
			return
		}
	}
	return
}
