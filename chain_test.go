package fp

import (
	"fmt"
	"reflect"
	"testing"
)

func SelectEven(i int) bool {
	return i%2 == 0
}

func Square(i int) int {
	return i * i
}

func Add(i int, j int) int {
	return i + j
}

func TestSlice_Reduce(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	t.Logf("got = %v", Slice[int](data).Filter(SelectEven).Map(Square).Reduce(Add, 0))
}

func TestTransform(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	list := Transform(data, func(t int) string {
		return fmt.Sprintf("%d", t)
	})
	t.Logf("got = %v", list)
}

func TestTransformAsync(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	list := TransformAsync(data, func(t int) string {
		return fmt.Sprintf("%d", t)
	})
	t.Logf("got = %v", list)
}

func TestWrap(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	Wrap(data).Map(func(v int) int {
		fmt.Println(v)
		return v
	})
	Wrap(data).Monad(func(v int) {
		fmt.Println(v)
	})
}

func TestSlice_MonadAsync(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	Wrap(data).MonadAsync(func(v int) {
		fmt.Println(v)
	})
}

func TestSlice_Len_IsEmpty(t *testing.T) {
	if Wrap([]int{}).Len() != 0 {
		t.Fatal("empty Len")
	}
	if !Wrap([]int{}).IsEmpty() {
		t.Fatal("empty IsEmpty")
	}
	if Wrap([]int{1, 2}).IsEmpty() {
		t.Fatal("non-empty IsEmpty")
	}
}

func TestSlice_First_Last(t *testing.T) {
	if v, ok := Wrap([]int{}).First(); ok || v != 0 {
		t.Fatalf("empty First got %v %v", v, ok)
	}
	if v, ok := Wrap([]int{1, 2, 3}).First(); !ok || v != 1 {
		t.Fatalf("First got %v %v", v, ok)
	}
	if v, ok := Wrap([]int{}).Last(); ok || v != 0 {
		t.Fatalf("empty Last got %v %v", v, ok)
	}
	if v, ok := Wrap([]int{1, 2, 3}).Last(); !ok || v != 3 {
		t.Fatalf("Last got %v %v", v, ok)
	}
}

func TestSlice_Any_All_Count(t *testing.T) {
	s := Wrap([]int{1, 2, 3, 4})
	if !s.Any(SelectEven) {
		t.Fatal("Any even")
	}
	if s.All(SelectEven) {
		t.Fatal("All even false")
	}
	if !Wrap([]int{}).All(SelectEven) {
		t.Fatal("empty All true")
	}
	if got := s.Count(SelectEven); got != 2 {
		t.Fatalf("Count got %d", got)
	}
}

func TestSlice_Find_FindIndex(t *testing.T) {
	s := Wrap([]int{1, 2, 3, 4})
	if v, ok := s.Find(SelectEven); !ok || v != 2 {
		t.Fatalf("Find got %v %v", v, ok)
	}
	if _, ok := s.Find(func(i int) bool { return i > 99 }); ok {
		t.Fatal("Find should miss")
	}
	if i := s.FindIndex(SelectEven); i != 1 {
		t.Fatalf("FindIndex got %d", i)
	}
	if i := s.FindIndex(func(i int) bool { return i > 99 }); i != -1 {
		t.Fatalf("FindIndex miss got %d", i)
	}
}

func TestSlice_Contains_IndexOf(t *testing.T) {
	s := Wrap([]int{10, 20, 30})
	if !s.Contains(20, Equal[int]) {
		t.Fatal("Contains 20")
	}
	if s.Contains(99, Equal[int]) {
		t.Fatal("Contains 99")
	}
	if i := s.IndexOf(30, Equal[int]); i != 2 {
		t.Fatalf("IndexOf 30 got %d", i)
	}
	if i := s.IndexOf(99, Equal[int]); i != -1 {
		t.Fatalf("IndexOf 99 got %d", i)
	}
}

func TestSlice_Reverse(t *testing.T) {
	got := Wrap([]int{1, 2, 3}).Reverse()
	if !reflect.DeepEqual([]int(got), []int{3, 2, 1}) {
		t.Fatalf("Reverse got %v", got)
	}
	if got := Wrap([]int{}).Reverse(); len(got) != 0 {
		t.Fatalf("Reverse empty got %v", got)
	}
}

func TestSlice_Take_Drop(t *testing.T) {
	s := Wrap([]int{1, 2, 3, 4, 5})
	if got := s.Take(2); !reflect.DeepEqual([]int(got), []int{1, 2}) {
		t.Fatalf("Take 2 got %v", got)
	}
	if got := s.Take(0); len(got) != 0 {
		t.Fatalf("Take 0 got %v", got)
	}
	if got := s.Take(100); !reflect.DeepEqual([]int(got), []int{1, 2, 3, 4, 5}) {
		t.Fatalf("Take 100 got %v", got)
	}
	if got := s.Drop(2); !reflect.DeepEqual([]int(got), []int{3, 4, 5}) {
		t.Fatalf("Drop 2 got %v", got)
	}
	if got := s.Drop(0); !reflect.DeepEqual([]int(got), []int{1, 2, 3, 4, 5}) {
		t.Fatalf("Drop 0 got %v", got)
	}
	if got := s.Drop(100); len(got) != 0 {
		t.Fatalf("Drop 100 got %v", got)
	}
}

func TestSlice_TakeWhile_DropWhile(t *testing.T) {
	s := Wrap([]int{2, 4, 5, 6, 8})
	if got := s.TakeWhile(SelectEven); !reflect.DeepEqual([]int(got), []int{2, 4}) {
		t.Fatalf("TakeWhile got %v", got)
	}
	if got := s.DropWhile(SelectEven); !reflect.DeepEqual([]int(got), []int{5, 6, 8}) {
		t.Fatalf("DropWhile got %v", got)
	}
	if got := Wrap([]int{1, 3}).TakeWhile(SelectEven); len(got) != 0 {
		t.Fatalf("TakeWhile none got %v", got)
	}
	if got := Wrap([]int{2, 4}).DropWhile(SelectEven); len(got) != 0 {
		t.Fatalf("DropWhile all got %v", got)
	}
}

func TestSlice_Chunk(t *testing.T) {
	got := Wrap([]int{1, 2, 3, 4, 5}).Chunk(2)
	want := []Slice[int]{{1, 2}, {3, 4}, {5}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Chunk got %v", got)
	}
	if got := Wrap([]int{1, 2}).Chunk(0); len(got) != 0 {
		t.Fatalf("Chunk size 0 got %v", got)
	}
	if got := Wrap([]int{}).Chunk(2); len(got) != 0 {
		t.Fatalf("Chunk empty got %v", got)
	}
}

func TestSlice_Concat(t *testing.T) {
	got := Wrap([]int{1, 2}).Concat([]int{3, 4})
	if !reflect.DeepEqual([]int(got), []int{1, 2, 3, 4}) {
		t.Fatalf("Concat got %v", got)
	}
}

func TestSlice_Partition(t *testing.T) {
	even, odd := Wrap([]int{1, 2, 3, 4, 5}).Partition(SelectEven)
	if !reflect.DeepEqual([]int(even), []int{2, 4}) {
		t.Fatalf("Partition even got %v", even)
	}
	if !reflect.DeepEqual([]int(odd), []int{1, 3, 5}) {
		t.Fatalf("Partition odd got %v", odd)
	}
}

func TestSlice_FilterWithKeyFunc(t *testing.T) {
	result := Wrap([]int{1, 2, 3, 4, 5, 1}).FilterWithKeyFunc(func(i int) string {
		return fmt.Sprintf("%d", i)
	})
	t.Logf("%+v", result)
	if !reflect.DeepEqual(result, Wrap([]int{1, 2, 3, 4, 5})) {
		t.Fatalf("FilterWithKeyFunc got %v", result)
	}
}

func TestSlice_FilterInSlice(t *testing.T) {
	result := Wrap([]int{1, 2, 3, 4, 5}).FilterInSlice([]int{1, 2, 3}, func(i1 int, i2 int) bool {
		return i1 == i2
	})
	t.Logf("%+v", result)
	if !reflect.DeepEqual(result, Wrap([]int{1, 2, 3})) {
		t.Fatalf("FilterWithKeyFunc got %v", result)
	}
}

func TestSlice_RemoveDuplicate(t *testing.T) {
	result := Wrap([]int{1, 2, 3, 4, 5, 6, 3, 3, 4, 2, 1, 1}).RemoveDuplicate(func(a int, b int) bool {
		return a == b
	})
	t.Logf("%+v", result)
	if !reflect.DeepEqual(result, Wrap([]int{1, 2, 3, 4, 5, 6})) {
		t.Fatalf("FilterWithKeyFunc got %v", result)
	}
}
