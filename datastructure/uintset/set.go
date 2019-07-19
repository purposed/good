package uintset

// A UintSet is a set of uint32s
type UintSet map[uint32]struct{}

// New returns a new uint32 set.
func New() UintSet {
	return make(map[uint32]struct{})
}

// FromValues initializes a uint32 set with values.
func FromValues(values []uint32) UintSet {
	x := New()
	for _, v := range values {
		x.Add(v)
	}
	return x
}

// Contains tests if the item is in the set.
func (s UintSet) Contains(itm uint32) bool {
	_, ok := s[itm]
	return ok
}

// Add adds an item to the set.
func (s UintSet) Add(itm uint32) {
	s[itm] = struct{}{}
}

// Remove removes an item from the set.
func (s UintSet) Remove(itm uint32) {
	delete(s, itm)
}

// Union updates current uint32 set to reflect the union between sets.
func (s UintSet) Union(other UintSet) UintSet {
	for k := range other {
		s.Add(k)
	}
	return s
}

// Intersection updates the current set to reflect the intersection between sets.
func (s UintSet) Intersection(other UintSet) UintSet {
	for k := range s {
		if !other.Contains(k) {
			s.Remove(k)
		}
	}
	return s
}

// Values returns the values in the set.
func (s UintSet) Values() []uint32 {
	vals := make([]uint32, len(s))

	i := 0
	for k := range s {
		vals[i] = k
		i++
	}

	return vals
}

// Equals checks for set equality.
func (s UintSet) Equals(other UintSet) bool {
	for item := range s {
		if !other.Contains(item) {
			return false
		}
	}

	for item := range other {
		if !s.Contains(item) {
			return false
		}
	}

	return true
}
