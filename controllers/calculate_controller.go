package controllers

import (
	"context"
	"net/http"
	"rest/configs"
	"rest/models"
	"rest/responses"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection = configs.GetCollectionName()
var timeEntryCollection = configs.GetCollection(configs.DB, collection)

func Calculate(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var data []models.Data
	defer cancel()

	result, err := timeEntryCollection.Aggregate(ctx, bson.A{
		bson.D{
			{Key: "$match",
				Value: bson.D{
					{Key: "startDateTime", Value: bson.D{{Key: "$gte", Value: "1567276200000"}}},
					{Key: "endDateTime", Value: bson.D{{Key: "$lte", Value: "1569868199000"}}},
				},
			},
		},
		bson.D{
			{Key: "$group",
				Value: bson.D{
					{Key: "_id", Value: "$resourceID"},
					{Key: "totalTime", Value: bson.D{{Key: "$sum", Value: bson.D{{Key: "$toDecimal", Value: "$timeSpent"}}}}},
				},
			},
		},
		bson.D{{Key: "$sort", Value: bson.D{{Key: "_id", Value: 1}}}},
	})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.DataResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	defer func(result *mongo.Cursor, ctx context.Context) {
		err := result.Close(ctx)
		if err != nil {

		}
	}(result, ctx)

	for result.Next(ctx) {
		var singleData models.Data

		if err := result.Decode(&singleData); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.DataResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		data = append(data, singleData)
	}

	return c.Status(http.StatusOK).JSON(
		responses.DataResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": data}},
	)
}

func CalculateWithoutAgg(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Find().SetProjection(bson.D{
		{Key: "_id", Value: 0},
		{Key: "resourceID", Value: 1},
		{Key: "timeSpent", Value: 1},
		{Key: "startDateTime", Value: 1},
		{Key: "endDateTime", Value: 1},
	})

	results, err := timeEntryCollection.Find(ctx, bson.D{
		{Key: "startDateTime", Value: bson.D{{Key: "$gte", Value: "1567276200000"}}},
		{Key: "endDateTime", Value: bson.D{{Key: "$lte", Value: "1569868199000"}}},
	}, opts)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.DataResponse{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data: &fiber.Map{
				"data": err.Error(),
			},
		})
	}

	defer func(results *mongo.Cursor, ctx context.Context) {
		err := results.Close(ctx)
		if err != nil {
		}
	}(results, ctx)

	dataWithTime := make(map[string]string)
	for results.Next(ctx) {
		var singleObject models.NewData
		if err = results.Decode(&singleObject); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.DataResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data: &fiber.Map{
					"data": err.Error(),
				},
			})
		}

		resourceId := singleObject.ResourceId
		timeSpent := singleObject.TimeSpent
		val, present := dataWithTime[resourceId]

		if !present {
			dataWithTime[resourceId] = timeSpent
		} else {
			newTime, _ := strconv.Atoi(timeSpent)
			existingTimeSpent, _ := strconv.Atoi(val)
			dataWithTime[resourceId] = strconv.Itoa(newTime + existingTimeSpent)
		}
	}

	return c.Status(http.StatusOK).JSON(
		responses.DataResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data: &fiber.Map{
				"data": dataWithTime,
			},
		})
}
