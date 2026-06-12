package fp

import (
	"reflect"
	"sort"
	"testing"
)

func TestMapKeys_MapValues(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	keys := MapKeys(m)
	values := MapValues(m)
	sort.Strings([]string(keys))
	sort.Ints([]int(values))
	if !reflect.DeepEqual([]string(keys), []string{"a", "b", "c"}) {
		t.Fatalf("MapKeys got %v", keys)
	}
	if !reflect.DeepEqual([]int(values), []int{1, 2, 3}) {
		t.Fatalf("MapValues got %v", values)
	}
}

func TestMapToSlice(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	got := MapToSlice[string, int, string](m, func(k string, v int) string {
		return k + ":" + string(rune('0'+v))
	})
	sort.Strings([]string(got))
	if !reflect.DeepEqual([]string(got), []string{"a:1", "b:2"}) {
		t.Fatalf("MapToSlice got %v", got)
	}
}

func TestMergeMaps(t *testing.T) {
	got := MergeMaps(
		map[string]int{"a": 1, "b": 2},
		map[string]int{"b": 20, "c": 3},
	)
	if !reflect.DeepEqual(got, map[string]int{"a": 1, "b": 20, "c": 3}) {
		t.Fatalf("MergeMaps got %v", got)
	}
}

func TestMapFilter(t *testing.T) {
	got := MapFilter(map[string]int{"a": 1, "b": 2, "c": 3}, func(_ string, v int) bool {
		return v >= 2
	})
	if !reflect.DeepEqual(got, map[string]int{"b": 2, "c": 3}) {
		t.Fatalf("MapFilter got %v", got)
	}
}
