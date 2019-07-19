package uintset_test

import (
	"testing"

	"github.com/purposed/good/datastructure/uintset"
)

func Test_NewUintSet(t *testing.T) {
	s := uintset.New()

	if s == nil {
		t.Error("uint32 was initialized to nil")
	}
}

func Test_UintSetFromValues(t *testing.T) {
	values := []uint32{1, 2, 3, 4}
	s2 := uintset.FromValues(values)

	for _, v := range values {
		if !s2.Contains(v) {
			t.Errorf("set does not contain value: %d", v)
			return
		}
	}
}

func TestUintSet_Contains(t *testing.T) {
	type args struct {
		itm uint32
	}
	tests := []struct {
		name string
		s    uintset.UintSet
		args args
		want bool
	}{
		{"empty set", uintset.FromValues(nil), args{1}, false},
		{"full set, item present", uintset.FromValues([]uint32{1}), args{1}, true},
		{"full set, item not present", uintset.FromValues([]uint32{1}), args{2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Contains(tt.args.itm); got != tt.want {
				t.Errorf("UintSet.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUintSet_Add(t *testing.T) {
	set := uintset.New()

	if set.Contains(1) {
		t.Error("set should begin empty")
	}

	set.Add(1)

	if !set.Contains(1) {
		t.Error("Add() did not add element to set")
	}

	set.Add(1)

	if len(set.Values()) > 1 {
		t.Error("item was added twice")
	}
}

func TestUintSet_Remove(t *testing.T) {
	set := uintset.FromValues([]uint32{1})
	set.Remove(1)

	if set.Contains(1) {
		t.Error("Remove() did not remove element from set")
	}
}

func TestUintSet_Union(t *testing.T) {
	type testCase struct {
		name string
		setA []uint32
		setB []uint32

		unionResult []uint32
	}

	cases := []testCase{
		{"empty sets", nil, nil, nil},
		{"a empty", nil, []uint32{1, 2}, []uint32{1, 2}},
		{"b empty", []uint32{1, 2}, nil, []uint32{1, 2}},
		{"shared", []uint32{1}, []uint32{2}, []uint32{1, 2}},
	}

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			setA := uintset.FromValues(tCase.setA)
			setB := uintset.FromValues(tCase.setB)

			expectedSet := uintset.FromValues(tCase.unionResult)

			setA.Union(setB)

			for _, itm := range expectedSet.Values() {
				if !setA.Contains(itm) {
					t.Errorf("expected set A to contain item: %d", itm)
					return
				}
			}

			for _, itm := range setA.Values() {
				if !expectedSet.Contains(itm) {
					t.Errorf("extra item in set A: %d", itm)
					return
				}
			}
		})
	}
}

func TestUintSet_Intersection(t *testing.T) {
	type testCase struct {
		name string
		setA []uint32
		setB []uint32

		unionResult []uint32
	}

	cases := []testCase{
		{"empty sets", nil, nil, nil},
		{"a empty", nil, []uint32{1, 2}, nil},
		{"b empty", []uint32{1, 2}, nil, nil},
		{"shared", []uint32{1}, []uint32{2}, nil},
		{"common items", []uint32{1, 3}, []uint32{2, 3}, []uint32{3}},
	}

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			setA := uintset.New()
			setB := uintset.New()

			expectedSet := uintset.New()

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
					t.Errorf("expected set A to contain item: %d", itm)
					return
				}
			}

			for _, itm := range setA.Values() {
				if !expectedSet.Contains(itm) {
					t.Errorf("extra item in set A: %d", itm)
					return
				}
			}
		})
	}
}

func TestUintSet_Values(t *testing.T) {
	type testCase struct {
		name     string
		setItems []uint32
	}

	cases := []testCase{
		{"empty set", nil},
		{"one-item set", []uint32{1}},
		{"multi set", []uint32{1, 2}},
	}

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			s := uintset.New()

			expected := uintset.New()
			actual := uintset.New()

			for _, k := range tCase.setItems {
				expected.Add(k)
				s.Add(k)
			}

			for _, k := range s.Values() {
				actual.Add(k)
			}

			for _, k := range actual.Values() {
				if !expected.Contains(k) {
					t.Errorf("unexpected item in result set: %d", k)
				}
			}

			for _, k := range expected.Values() {
				if !actual.Contains(k) {
					t.Errorf("missing item in result set: %d", k)
				}
			}
		})
	}
}

func TestUintSet_Equals(t *testing.T) {
	type testCase struct {
		name    string
		valuesA []uint32
		valuesB []uint32

		equal bool
	}

	cases := []testCase{
		{"empty sets", nil, nil, true},
		{"one empty", []uint32{1}, nil, false},
		{"both full - true", []uint32{2, 1}, []uint32{1, 2}, true},
		{"both full - false", []uint32{1}, []uint32{1, 2}, false},
	}

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			a := uintset.FromValues(tCase.valuesA)
			b := uintset.FromValues(tCase.valuesB)

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
