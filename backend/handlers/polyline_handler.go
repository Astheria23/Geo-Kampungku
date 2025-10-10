package handlers

import (
	"context"
	"net/http"
	"time"

	"posttest/backend/database"
	"posttest/backend/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const defaultTimeout = 10 * time.Second

// CreatePolyline handles POST /api/polylines
func CreatePolyline(c *fiber.Ctx) error {
	var input models.PolylineInput
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid JSON payload")
	}
	if err := input.Geometry.Validate(); err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}
	if input.Name == "" {
		return fiber.NewError(http.StatusBadRequest, "name is required")
	}
	if input.Region == "" {
		return fiber.NewError(http.StatusBadRequest, "region is required")
	}

	polyline := models.Polyline{}
	polyline.Apply(input)
	polyline.ID = primitive.NewObjectID()
	polyline.CreatedAt = time.Now().UTC()
	polyline.UpdatedAt = polyline.CreatedAt

	collection, err := database.Collection()
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	ctx, cancel := context.WithTimeout(c.Context(), defaultTimeout)
	defer cancel()

	if _, err = collection.InsertOne(ctx, polyline); err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.Status(http.StatusCreated).JSON(polyline)
}

// GetPolylines handles GET /api/polylines
func GetPolylines(c *fiber.Ctx) error {
	collection, err := database.Collection()
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	ctx, cancel := context.WithTimeout(c.Context(), defaultTimeout)
	defer cancel()

	filter := bson.M{}
	if region := c.Query("region"); region != "" {
		filter["region"] = region
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}
	defer cursor.Close(ctx)

	var polylines []models.Polyline
	if err = cursor.All(ctx, &polylines); err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(polylines)
}

// GetPolyline handles GET /api/polylines/:id
func GetPolyline(c *fiber.Ctx) error {
	idParam := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid id format")
	}

	collection, err := database.Collection()
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	ctx, cancel := context.WithTimeout(c.Context(), defaultTimeout)
	defer cancel()

	var polyline models.Polyline
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&polyline)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return fiber.NewError(http.StatusNotFound, "polyline not found")
		}
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(polyline)
}

// UpdatePolyline handles PUT /api/polylines/:id
func UpdatePolyline(c *fiber.Ctx) error {
	idParam := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid id format")
	}

	var input models.PolylineInput
	if err = c.BodyParser(&input); err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid JSON payload")
	}
	if err = input.Geometry.Validate(); err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	collection, err := database.Collection()
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	ctx, cancel := context.WithTimeout(c.Context(), defaultTimeout)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"name":        input.Name,
			"description": input.Description,
			"region":      input.Region,
			"geometry":    input.Geometry,
			"updatedAt":   time.Now().UTC(),
		},
	}

	result, err := collection.UpdateByID(ctx, objID, update)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}
	if result.MatchedCount == 0 {
		return fiber.NewError(http.StatusNotFound, "polyline not found")
	}

	var updated models.Polyline
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&updated)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(updated)
}

// DeletePolyline handles DELETE /api/polylines/:id
func DeletePolyline(c *fiber.Ctx) error {
	idParam := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid id format")
	}

	collection, err := database.Collection()
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	ctx, cancel := context.WithTimeout(c.Context(), defaultTimeout)
	defer cancel()

	result, err := collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}
	if result.DeletedCount == 0 {
		return fiber.NewError(http.StatusNotFound, "polyline not found")
	}

	return c.SendStatus(http.StatusNoContent)
}
