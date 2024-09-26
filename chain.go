package fp

import "sync"

// Slice 是一个泛型切片类型
type Slice[T any] []T

// Filter 函数根据条件过滤切片中的元素
func (s Slice[T]) Filter(condition func(T) bool) Slice[T] {
	var result Slice[T]
	for _, v := range s {
		if condition(v) {
			result = append(result, v)
		}
	}
	return result
}

// Map 函数对切片中的每个元素应用变换函数
func (s Slice[T]) Map(transform func(T) T) Slice[T] {
	var result Slice[T]
	for _, v := range s {
		result = append(result, transform(v))
	}
	return result
}

// Reduce 函数将切片归约为一个值
func (s Slice[T]) Reduce(accumulator func(T, T) T, initial T) T {
	result := initial
	for _, v := range s {
		result = accumulator(result, v)
	}
	return result
}

// Transform 函数将切片中的每一个值转成新类型
func Transform[T any, B any](s Slice[T], transform func(T) B) Slice[B] {
	var result Slice[B]
	for _, v := range s {
		result = append(result, transform(v))
	}
	return result
}

func Wrap[T any](s Slice[T]) Slice[T] {
	return s
}

func (s Slice[T]) Monad(f func(T)) Slice[T] {
	for _, v := range s {
		f(v)
	}
	return s
}

func (s Slice[T]) MonadAsync(f func(T)) Slice[T] {
	var wg sync.WaitGroup
	for _, v := range s {
		wg.Add(1)
		go func(v T) {
			defer wg.Done()
			f(v)
		}(v)
	}
	wg.Wait()
	return s
}
