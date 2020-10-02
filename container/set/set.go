package set

// Set is ...
type Set map[interface{}]struct{}

// Has func checks whether s has item or not
func (s Set) Has(item interface{}) bool {
	_, exists := s[item]
	return exists
}

func (s Set) Insert(item interface{}) {
	s[item] = struct{}{}
}

func (s Set) Delete(item interface{}) {
	delete(s, item)
}
