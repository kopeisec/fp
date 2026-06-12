package fp

import (
	"reflect"
	"sort"
	"strconv"
	"testing"
)

func TestFlatMap(t *testing.T) {
	got := FlatMap[int, int](Slice[int]{1, 2, 3}, func(i int) Slice[int] {
		return Slice[int]{i, i * 10}
	})
	if !reflect.DeepEqual([]int(got), []int{1, 10, 2, 20, 3, 30}) {
		t.Fatalf("FlatMap got %v", got)
	}
}

func TestFlatten(t *testing.T) {
	got := Flatten(Slice[Slice[int]]{{1, 2}, {3}, {4, 5}})
	if !reflect.DeepEqual([]int(got), []int{1, 2, 3, 4, 5}) {
		t.Fatalf("Flatten got %v", got)
	}
}

func TestGroupBy(t *testing.T) {
	got := GroupBy[int, string](Slice[int]{1, 2, 3, 4, 5}, func(i int) string {
		if i%2 == 0 {
			return "even"
		}
		return "odd"
	})
	if !reflect.DeepEqual([]int(got["even"]), []int{2, 4}) {
		t.Fatalf("GroupBy even got %v", got["even"])
	}
	if !reflect.DeepEqual([]int(got["odd"]), []int{1, 3, 5}) {
		t.Fatalf("GroupBy odd got %v", got["odd"])
	}
}

func TestZip_Unzip(t *testing.T) {
	zipped := Zip[int, string](Slice[int]{1, 2, 3}, Slice[string]{"a", "b"})
	if len(zipped) != 2 || zipped[0].First != 1 || zipped[0].Second != "a" {
		t.Fatalf("Zip got %v", zipped)
	}
	as, bs := Unzip(zipped)
	if !reflect.DeepEqual([]int(as), []int{1, 2}) {
		t.Fatalf("Unzip as got %v", as)
	}
	if !reflect.DeepEqual([]string(bs), []string{"a", "b"}) {
		t.Fatalf("Unzip bs got %v", bs)
	}
}

func TestReduceTo(t *testing.T) {
	got := ReduceTo[int, string](Slice[int]{1, 2, 3}, func(acc string, v int) string {
		return acc + strconv.Itoa(v)
	}, "")
	if got != "123" {
		t.Fatalf("ReduceTo got %v", got)
	}
}

func TestGroupBy_KeysStable(t *testing.T) {
	m := GroupBy[int, int](Slice[int]{1, 2, 3, 4}, func(i int) int { return i % 2 })
	keys := MapKeys(m)
	asInts := []int(keys)
	sort.Ints(asInts)
	if !reflect.DeepEqual(asInts, []int{0, 1}) {
		t.Fatalf("GroupBy keys got %v", asInts)
	}
}
