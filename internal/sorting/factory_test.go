package sorting

import (
	"testing"
)

func TestNewPersonSorter(t *testing.T) {
	tests := []struct {
		name    string
		field   string
		wantNil bool
	}{
		{
			name:    "name sorter",
			field:   "name",
			wantNil: false,
		},
		{
			name:    "created sorter",
			field:   "created",
			wantNil: false,
		},
		{
			name:    "mass sorter",
			field:   "mass",
			wantNil: false,
		},
		{
			name:    "unknown field returns nil",
			field:   "unknown",
			wantNil: true,
		},
		{
			name:    "empty field returns nil",
			field:   "",
			wantNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sorter := NewPersonSorter(tt.field)

			if tt.wantNil {
				if sorter != nil {
					t.Errorf("NewPersonSorter(%q) = %v, want nil", tt.field, sorter)
				}
			} else {
				if sorter == nil {
					t.Errorf("NewPersonSorter(%q) = nil, want non-nil", tt.field)
				}
			}
		})
	}
}

func TestNewPlanetSorter(t *testing.T) {
	tests := []struct {
		name    string
		field   string
		wantNil bool
	}{
		{
			name:    "name sorter",
			field:   "name",
			wantNil: false,
		},
		{
			name:    "created sorter",
			field:   "created",
			wantNil: false,
		},
		{
			name:    "mass not supported for planets",
			field:   "mass",
			wantNil: true,
		},
		{
			name:    "unknown field returns nil",
			field:   "unknown",
			wantNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sorter := NewPlanetSorter(tt.field)

			if tt.wantNil {
				if sorter != nil {
					t.Errorf("NewPlanetSorter(%q) = %v, want nil", tt.field, sorter)
				}
			} else {
				if sorter == nil {
					t.Errorf("NewPlanetSorter(%q) = nil, want non-nil", tt.field)
				}
			}
		})
	}
}
