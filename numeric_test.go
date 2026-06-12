package fp

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	if got := Sum(Slice[int]{1, 2, 3, 4}); got != 10 {
		t.Fatalf("Sum int got %d", got)
	}
	if got := Sum(Slice[float64]{1.5, 2.5}); got != 4.0 {
		t.Fatalf("Sum float got %v", got)
	}
	if got := Sum(Slice[int]{}); got != 0 {
		t.Fatalf("Sum empty got %d", got)
	}
}

func TestMin_Max(t *testing.T) {
	if v, ok := Min(Slice[int]{3, 1, 4, 1, 5}); !ok || v != 1 {
		t.Fatalf("Min got %v %v", v, ok)
	}
	if v, ok := Max(Slice[int]{3, 1, 4, 1, 5}); !ok || v != 5 {
		t.Fatalf("Max got %v %v", v, ok)
	}
	if _, ok := Min(Slice[int]{}); ok {
		t.Fatal("Min empty ok")
	}
	if _, ok := Max(Slice[int]{}); ok {
		t.Fatal("Max empty ok")
	}
}

func TestAvg(t *testing.T) {
	if v, ok := Avg(Slice[int]{2, 4, 6}); !ok || v != 4.0 {
		t.Fatalf("Avg got %v %v", v, ok)
	}
	if _, ok := Avg(Slice[int]{}); ok {
		t.Fatal("Avg empty ok")
	}
}

func TestSort(t *testing.T) {
	got := Sort(Slice[int]{3, 1, 4, 1, 5, 9, 2, 6})
	if !reflect.DeepEqual([]int(got), []int{1, 1, 2, 3, 4, 5, 6, 9}) {
		t.Fatalf("Sort got %v", got)
	}
}

func TestSortBy(t *testing.T) {
	type item struct {
		name string
		age  int
	}
	got := SortBy[item, int](Slice[item]{{"a", 3}, {"b", 1}, {"c", 2}}, func(i item) int {
		return i.age
	})
	want := []item{{"b", 1}, {"c", 2}, {"a", 3}}
	if !reflect.DeepEqual([]item(got), want) {
		t.Fatalf("SortBy got %v", got)
	}
}
