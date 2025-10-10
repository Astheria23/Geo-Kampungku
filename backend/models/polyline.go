package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Polyline represents a road segment stored as GeoJSON LineString.
type Polyline struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Region      string             `json:"region" bson:"region"`
	Geometry    LineString         `json:"geometry" bson:"geometry"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt" bson:"updatedAt"`
}

// PolylineInput represents incoming payload for create/update.
type PolylineInput struct {
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Region      string     `json:"region"`
	Geometry    LineString `json:"geometry"`
}

// Apply populates Polyline fields from input while managing timestamps.
func (p *Polyline) Apply(input PolylineInput) {
	p.Name = input.Name
	p.Description = input.Description
	p.Region = input.Region
	p.Geometry = input.Geometry
}
