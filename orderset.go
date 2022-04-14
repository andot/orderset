package go_orderset

type OrderSet[K comparable] struct {
	set   map[K]int
	order []K
}

func NewOrderSet[K comparable]() OrderSet[K] {
	return OrderSet[K]{set: make(map[K]int), order: make([]K, 0)}
}

func (s *OrderSet[K]) Add(element K) {
	if _, exist := s.set[element]; !exist {
		// Get index by array length
		index := len(s.order)
		s.set[element] = index
		s.order = append(s.order, element)
	}
}

func (s *OrderSet[K]) Del(element K) {
	if index, exist := s.set[element]; exist {
		// Get index and delete element
		delete(s.set, element)
		s.order = append(s.order[:index], s.order[index+1:]...)
		s.reIndex()
	}
}

func (s *OrderSet[K]) Contains(element K) bool {
	_, exist := s.set[element]
	return exist
}

func (s *OrderSet[K]) ToSlice() []K {
	return s.order
}

func (s *OrderSet[K]) Count() int {
	return len(s.order)
}

func (s *OrderSet[K]) reIndex() {
	for i, k := range s.order {
		s.set[k] = i
	}
}
