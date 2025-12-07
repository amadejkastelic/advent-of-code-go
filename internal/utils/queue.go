package utils

type Item[T any] struct {
	value T
	index int
}

type Queue[T any] []*Item[T]

func (pq Queue[T]) Len() int { return len(pq) }

func (pq Queue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *Queue[T]) Push(x T) {
	n := len(*pq)
	item := &Item[T]{value: x, index: n}
	*pq = append(*pq, item)
}

func (pq *Queue[T]) Pop() T {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item.value
}
