package stringset

// A StringSet is a set of strings
type StringSet map[string]struct{}

// New returns a new string set.
func New() StringSet {
	return make(map[string]struct{})
}

// FromValues initializes a string set with values.
func FromValues(values []string) StringSet {
	x := New()
	for _, v := range values {
		x.Add(v)
	}
	return x
}

// Contains tests if the item is in the set.
func (s StringSet) Contains(itm string) bool {
	_, ok := s[itm]
	return ok
}

// Add adds an item to the set.
func (s StringSet) Add(itm string) {
	s[itm] = struct{}{}
}

// Remove removes an item from the set.
func (s StringSet) Remove(itm string) {
	delete(s, itm)
}

// Union updates current string set to reflect the union between sets.
func (s StringSet) Union(other StringSet) StringSet {
	for k := range other {
		s.Add(k)
	}
	return s
}

// Intersection updates the current set to reflect the intersection between sets.
func (s StringSet) Intersection(other StringSet) StringSet {
	for k := range s {
		if !other.Contains(k) {
			s.Remove(k)
		}
	}
	return s
}

// Values returns the values in the set.
func (s StringSet) Values() []string {
	vals := make([]string, len(s))

	i := 0
	for k := range s {
		vals[i] = k
		i++
	}

	return vals
}

// Equals checks for set equality.
func (s StringSet) Equals(other StringSet) bool {
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
