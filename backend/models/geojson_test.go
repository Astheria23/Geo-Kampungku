package models

import "testing"

func TestLineStringValidate(t *testing.T) {
	cases := []struct {
		name      string
		geom      LineString
		wantError bool
	}{
		{
			name: "valid linestring",
			geom: LineString{
				Type:        "LineString",
				Coordinates: [][]float64{{110.0, -7.0}, {110.1, -7.1}},
			},
			wantError: false,
		},
		{
			name: "invalid type",
			geom: LineString{
				Type:        "Polygon",
				Coordinates: [][]float64{{110.0, -7.0}, {110.1, -7.1}},
			},
			wantError: true,
		},
		{
			name: "insufficient points",
			geom: LineString{
				Type:        "LineString",
				Coordinates: [][]float64{{110.0, -7.0}},
			},
			wantError: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.geom.Validate()
			if tc.wantError && err == nil {
				t.Fatalf("expected error, got nil")
			}
			if !tc.wantError && err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
		})
	}
}
