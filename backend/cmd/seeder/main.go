package main

import (
	"context"
	"log"
	"time"

	"posttest/backend/config"
	"posttest/backend/database"
	"posttest/backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	config.LoadEnv()

	collection, err := database.Collection()
	if err != nil {
		log.Fatalf("seeder: gagal mendapatkan koleksi: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	seeds := polylineSeeds()

	var inserted, updated, skipped int

	for _, seed := range seeds {
		if err := seed.Geometry.Validate(); err != nil {
			log.Printf("seeder: melewati %s (%s) karena data tidak valid: %v", seed.Name, seed.Region, err)
			skipped++
			continue
		}

		now := time.Now().UTC()

		update := bson.M{
			"$set": bson.M{
				"name":        seed.Name,
				"description": seed.Description,
				"region":      seed.Region,
				"geometry":    seed.Geometry,
				"updatedAt":   now,
			},
			"$setOnInsert": bson.M{
				"_id":       primitive.NewObjectID(),
				"createdAt": now,
			},
		}

		result, err := collection.UpdateOne(
			ctx,
			bson.M{"name": seed.Name, "region": seed.Region},
			update,
			options.Update().SetUpsert(true),
		)
		if err != nil {
			log.Printf("seeder: gagal menyimpan %s (%s): %v", seed.Name, seed.Region, err)
			skipped++
			continue
		}

		if result.UpsertedCount > 0 {
			inserted++
		} else if result.ModifiedCount > 0 {
			updated++
		}
	}

	log.Printf("seeder selesai. inserted=%d updated=%d skipped=%d total=%d", inserted, updated, skipped, len(seeds))
}

func polylineSeeds() []models.Polyline {
	return []models.Polyline{
		{
			Name:        "Jalan Raya Utama",
			Description: "Ruas utama menuju kantor kelurahan",
			Region:      "Kelurahan Pulosari",
			Geometry: models.LineString{
				Type: "LineString",
				Coordinates: [][]float64{
					{110.43012, -7.12345},
					{110.43157, -7.12410},
					{110.43302, -7.12485},
				},
			},
		},
		{
			Name:        "Jalan Pendidikan",
			Description: "Akses menuju sekolah dan kampus",
			Region:      "Kelurahan Pulosari",
			Geometry: models.LineString{
				Type: "LineString",
				Coordinates: [][]float64{
					{110.42980, -7.12210},
					{110.43065, -7.12135},
					{110.43150, -7.12050},
				},
			},
		},
		{
			Name:        "Lingkar Kampung",
			Description: "Jalan lingkar yang menghubungkan beberapa RT",
			Region:      "Dusun Sumber Rejeki",
			Geometry: models.LineString{
				Type: "LineString",
				Coordinates: [][]float64{
					{110.43510, -7.12600},
					{110.43430, -7.12690},
					{110.43340, -7.12750},
					{110.43260, -7.12810},
				},
			},
		},
	}
}
