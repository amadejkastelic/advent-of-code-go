package sliceutils

func Map[T any, U any](input []T, mapper func(T) U) []U {
	result := make([]U, len(input))
	for i, v := range input {
		result[i] = mapper(v)
	}
	return result
}

func MapWithError[T any, U any](input []T, mapper func(T) (U, error)) ([]U, error) {
	result := make([]U, len(input))
	for i, v := range input {
		mappedValue, err := mapper(v)
		if err != nil {
			return nil, err
		}
		result[i] = mappedValue
	}
	return result, nil
}

func Filter[T any](input []T, predicate func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range input {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func FilterWithError[T any](input []T, predicate func(T) (bool, error)) ([]T, error) {
	result := make([]T, 0)
	for _, v := range input {
		match, err := predicate(v)
		if err != nil {
			return nil, err
		}
		if match {
			result = append(result, v)
		}
	}
	return result, nil
}

func ToMap[T comparable, U any](input []T, valueFunc func(T) U) map[T]U {
	result := make(map[T]U)
	for _, v := range input {
		result[v] = valueFunc(v)
	}
	return result
}

func Sum[T ~int | ~int64 | ~float64](input []T) T {
	var sum T
	for _, v := range input {
		sum += v
	}
	return sum
}

func Reduce[T any, U any](input []T, reducer func(U, T) U, initial U) U {
	result := initial
	for _, v := range input {
		result = reducer(result, v)
	}
	return result
}

func All[T any](input []T, predicate func(T) bool) bool {
	for _, v := range input {
		if !predicate(v) {
			return false
		}
	}
	return true
}
