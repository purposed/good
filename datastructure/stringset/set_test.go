package stringset_test

import (
	"testing"

	"github.com/purposed/good/datastructure/stringset"
)

func Test_NewStringSet(t *testing.T) {
	s := stringset.New()

	if s == nil {
		t.Error("string was initialized to nil")
	}
}

func Test_StringSetFromValues(t *testing.T) {
	values := []string{"a", "b", "c", "d"}
	s2 := stringset.FromValues(values)

	for _, v := range values {
		if !s2.Contains(v) {
			t.Errorf("set does not contain value: %s", v)
			return
		}
	}
}

func TestStringSet_Contains(t *testing.T) {
	type args struct {
		itm string
	}
	tests := []struct {
		name string
		s    stringset.StringSet
		args args
		want bool
	}{
		{"empty set", stringset.FromValues(nil), args{"a"}, false},
		{"full set, item present", stringset.FromValues([]string{"a"}), args{"a"}, true},
		{"full set, item not present", stringset.FromValues([]string{"a"}), args{"b"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Contains(tt.args.itm); got != tt.want {
				t.Errorf("StringSet.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringSet_Add(t *testing.T) {
	set := stringset.New()

	if set.Contains("a") {
		t.Error("set should begin empty")
	}

	set.Add("a")

	if !set.Contains("a") {
		t.Error("Add() did not add element to set")
	}

	set.Add("a")

	if len(set.Values()) > 1 {
		t.Error("item was added twice")
	}
}

func TestStringSet_Remove(t *testing.T) {
	set := stringset.FromValues([]string{"a"})
	set.Remove("a")

	if set.Contains("a") {
		t.Error("Remove() did not remove element from set")
	}
}

func TestStringSet_Union(t *testing.T) {
	type testCase struct {
		name string
		setA []string
		setB []string

		unionResult []string
	}

	cases := []testCase{
		{"empty sets", nil, nil, nil},
		{"a empty", nil, []string{"a", "b"}, []string{"a", "b"}},
		{"b empty", []string{"a", "b"}, nil, []string{"a", "b"}},
		{"shared", []string{"a"}, []string{"b"}, []string{"a", "b"}},
	}

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			setA := stringset.FromValues(tCase.setA)
			setB := stringset.FromValues(tCase.setB)

			expectedSet := stringset.FromValues(tCase.unionResult)

			setA.Union(setB)

			for _, itm := range expectedSet.Values() {
				if !setA.Contains(itm) {
					t.Errorf("expected set A to contain item: %s", itm)
					return
				}
			}

			for _, itm := range setA.Values() {
				if !expectedSet.Contains(itm) {
					t.Errorf("extra item in set A: %s", itm)
					return
				}
			}
		})
	}
}

func TestStringSet_Intersection(t *testing.T) {
	type testCase struct {
		name string
		setA []string
		setB []string

		unionResult []string
	}

	cases := []testCase{
		{"empty sets", nil, nil, nil},
		{"a empty", nil, []string{"a", "b"}, nil},
		{"b empty", []string{"a", "b"}, nil, nil},
		{"shared", []string{"a"}, []string{"b"}, nil},
		{"common items", []string{"a", "c"}, []string{"b", "c"}, []string{"c"}},
	}

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			setA := stringset.New()
			setB := stringset.New()

			expectedSet := stringset.New()

			for _, k := range tCase.setA {
				setA.Add(k)
			}

			for _, k := range tCase.setB {
				setB.Add(k)
			}

			for _, k := range tCase.unionResult {
				expectedSet.Add(k)
			}

			setA.Intersection(setB)

			for _, itm := range expectedSet.Values() {
				if !setA.Contains(itm) {
					t.Errorf("expected set A to contain item: %s", itm)
					return
				}
			}

			for _, itm := range setA.Values() {
				if !expectedSet.Contains(itm) {
					t.Errorf("extra item in set A: %s", itm)
					return
				}
			}
		})
	}
}

func TestStringSet_Values(t *testing.T) {
	type testCase struct {
		name     string
		setItems []string
	}

	cases := []testCase{
		{"empty set", nil},
		{"one-item set", []string{"a"}},
		{"multi set", []string{"a", "b"}},
	}

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			s := stringset.New()

			expected := stringset.New()
			actual := stringset.New()

			for _, k := range tCase.setItems {
				expected.Add(k)
				s.Add(k)
			}

			for _, k := range s.Values() {
				actual.Add(k)
			}

			for _, k := range actual.Values() {
				if !expected.Contains(k) {
					t.Errorf("unexpected item in result set: %s", k)
				}
			}

			for _, k := range expected.Values() {
				if !actual.Contains(k) {
					t.Errorf("missing item in result set: %s", k)
				}
			}
		})
	}
}

func TestStringSet_Equals(t *testing.T) {
	type testCase struct {
		name    string
		valuesA []string
		valuesB []string

		equal bool
	}

	cases := []testCase{
		{"empty sets", nil, nil, true},
		{"one empty", []string{"hello"}, nil, false},
		{"both full - true", []string{"world", "hello"}, []string{"hello", "world"}, true},
		{"both full - false", []string{"hello"}, []string{"hello", "world"}, false},
	}

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			a := stringset.FromValues(tCase.valuesA)
			b := stringset.FromValues(tCase.valuesB)

			aa := a.Equals(a)
			bb := b.Equals(b)

			if !aa || !bb {
				t.Errorf("consistence error: set not equal to itself")
				return
			}

			aR := a.Equals(b)
			bR := b.Equals(a)

			if aR != bR {
				t.Error("commutativity error: a = b != b = a")
				return
			}

			if aR != tCase.equal {
				t.Errorf("expected equal=%v, got %v", tCase.equal, aR)
				return
			}
		})
	}
}
