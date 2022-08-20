package orderset

type OrderSet[K comparable] struct {
	set   map[K]int
	order []K
}

func NewOrderSet[K comparable]() *OrderSet[K] {
	return &OrderSet[K]{set: make(map[K]int), order: make([]K, 0)}
}

func (s *OrderSet[K]) Add(element K) *OrderSet[K] {
	if s == nil {
		s = NewOrderSet[K]()
	}
	if _, exist := s.set[element]; !exist {
		// Get index by array length
		index := len(s.order)
		s.set[element] = index
		s.order = append(s.order, element)
	}
	return s
}

func (s *OrderSet[K]) Del(element K) {
	if s == nil {
		return
	}
	if index, exist := s.set[element]; exist {
		// Get index and delete element
		delete(s.set, element)
		s.order = append(s.order[:index], s.order[index+1:]...)
		s.reIndex()
	}
}

func (s *OrderSet[K]) Contains(element K) bool {
	if s == nil {
		return false
	}
	if len(s.order) < 15 {
		for _, v := range s.order {
			if element == v {
				return true
			}
		}
		return false
	}
	_, exist := s.set[element]
	return exist
}

func (s *OrderSet[K]) Index(element K) int {
	if s == nil {
		return -1
	}
	if len(s.order) < 15 {
		for i, v := range s.order {
			if element == v {
				return i
			}
		}
		return -1
	}
	if i, exist := s.set[element]; exist {
		return i
	}
	return -1
}

func (s *OrderSet[K]) ToSlice() []K {
	if s == nil {
		return nil
	}
	return s.order
}

func (s *OrderSet[K]) Count() int {
	if s == nil {
		return 0
	}
	return len(s.order)
}

func (s *OrderSet[K]) reIndex() {
	for i, k := range s.order {
		s.set[k] = i
	}
}
