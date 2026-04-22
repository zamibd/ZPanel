package common

import (
	"sort"
	"testing"
)

func sortedUint(s []uint) []uint {
	cp := make([]uint, len(s))
	copy(cp, s)
	sort.Slice(cp, func(i, j int) bool { return cp[i] < cp[j] })
	return cp
}

func TestUnionUintArray(t *testing.T) {
	tests := []struct {
		name string
		a, b []uint
		want []uint
	}{
		{
			name: "disjoint slices",
			a:    []uint{1, 2},
			b:    []uint{3, 4},
			want: []uint{1, 2, 3, 4},
		},
		{
			name: "overlapping slices",
			a:    []uint{1, 2, 3},
			b:    []uint{2, 3, 4},
			want: []uint{1, 2, 3, 4},
		},
		{
			name: "identical slices",
			a:    []uint{1, 2},
			b:    []uint{1, 2},
			want: []uint{1, 2},
		},
		{
			name: "empty a",
			a:    []uint{},
			b:    []uint{1, 2},
			want: []uint{1, 2},
		},
		{
			name: "both empty",
			a:    []uint{},
			b:    []uint{},
			want: []uint{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortedUint(UnionUintArray(tt.a, tt.b))
			want := sortedUint(tt.want)
			if len(got) != len(want) {
				t.Fatalf("got len %d, want len %d: got=%v want=%v", len(got), len(want), got, want)
			}
			for i := range got {
				if got[i] != want[i] {
					t.Errorf("index %d: got %d, want %d", i, got[i], want[i])
				}
			}
		})
	}
}

func TestDiffUintArray(t *testing.T) {
	tests := []struct {
		name string
		a, b []uint
		want []uint // symmetric diff
	}{
		{
			name: "disjoint slices",
			a:    []uint{1, 2},
			b:    []uint{3, 4},
			want: []uint{1, 2, 3, 4},
		},
		{
			name: "overlapping slices",
			a:    []uint{1, 2, 3},
			b:    []uint{2, 3, 4},
			want: []uint{1, 4},
		},
		{
			name: "identical slices",
			a:    []uint{1, 2},
			b:    []uint{1, 2},
			want: []uint{},
		},
		{
			name: "empty a",
			a:    []uint{},
			b:    []uint{1, 2},
			want: []uint{1, 2},
		},
		{
			name: "both empty",
			a:    []uint{},
			b:    []uint{},
			want: []uint{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortedUint(DiffUintArray(tt.a, tt.b))
			want := sortedUint(tt.want)
			if len(got) != len(want) {
				t.Fatalf("got len %d, want len %d: got=%v want=%v", len(got), len(want), got, want)
			}
			for i := range got {
				if got[i] != want[i] {
					t.Errorf("index %d: got %d, want %d", i, got[i], want[i])
				}
			}
		})
	}
}
