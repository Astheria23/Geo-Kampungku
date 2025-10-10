package models

import (
	"errors"
)

// LineString represents a GeoJSON LineString geometry for roads.
type LineString struct {
	Type        string      `json:"type" bson:"type"`
	Coordinates [][]float64 `json:"coordinates" bson:"coordinates"`
}

// Validate ensures the LineString contains valid GeoJSON data.
func (ls *LineString) Validate() error {
	if ls == nil {
		return errors.New("linestring: geometry is required")
	}
	if ls.Type == "" {
		return errors.New("linestring: type is required")
	}
	if ls.Type != "LineString" {
		return errors.New("linestring: type must be LineString")
	}
	if len(ls.Coordinates) < 2 {
		return errors.New("linestring: at least two coordinates required")
	}
	for i, coord := range ls.Coordinates {
		if len(coord) != 2 {
			return errors.New("linestring: coordinates must be [lng, lat]")
		}
		if coord[0] < -180 || coord[0] > 180 {
			return errors.New("linestring: longitude out of range")
		}
		if coord[1] < -90 || coord[1] > 90 {
			return errors.New("linestring: latitude out of range")
		}
		_ = i
	}
	return nil
}
