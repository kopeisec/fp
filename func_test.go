package fp

import "testing"

func TestIdentity(t *testing.T) {
	if got := Identity(42); got != 42 {
		t.Fatalf("Identity got %d", got)
	}
	if got := Identity("hi"); got != "hi" {
		t.Fatalf("Identity got %s", got)
	}
}

func TestNot(t *testing.T) {
	notEven := Not(SelectEven)
	if notEven(2) {
		t.Fatal("Not even on 2")
	}
	if !notEven(3) {
		t.Fatal("Not even on 3")
	}
}

func TestCompose(t *testing.T) {
	addOne := func(i int) int { return i + 1 }
	double := func(i int) int { return i * 2 }
	doubleThenAdd := Compose(addOne, double)
	if got := doubleThenAdd(3); got != 7 {
		t.Fatalf("Compose got %d", got)
	}
}

func TestPipe(t *testing.T) {
	addOne := func(i int) int { return i + 1 }
	double := func(i int) int { return i * 2 }
	square := func(i int) int { return i * i }
	pipeline := Pipe(addOne, double, square)
	if got := pipeline(3); got != 64 {
		t.Fatalf("Pipe got %d", got)
	}
	if got := Pipe[int]()(5); got != 5 {
		t.Fatalf("Pipe empty got %d", got)
	}
}
