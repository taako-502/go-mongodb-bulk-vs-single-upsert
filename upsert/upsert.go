package upsert

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Upsert(collection *mongo.Collection, count int) error {
	id := primitive.NewObjectID()
	for i := range count {
		_, err := collection.UpdateOne(
			context.TODO(),
			bson.M{"_id": id},
			bson.M{
				"$set": bson.M{
					"_id":       id,
					"text":      "upsert",
					"createdAt": primitive.DateTime(time.Now().UnixNano() / int64(time.Millisecond)),
					"updatedAt": primitive.DateTime(time.Now().UnixNano() / int64(time.Millisecond)),
				},
			},
			options.Update().SetUpsert(true),
		)
		if err != nil {
			return fmt.Errorf("failed to upsert: %w, count: %d", err, i)
		}
	}

	return nil
}
