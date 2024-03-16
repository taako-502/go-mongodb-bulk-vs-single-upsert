package benchmark

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Cleanup(collection *mongo.Collection) error {
	_, err := collection.DeleteMany(context.TODO(), bson.M{})
	if err != nil {
		return fmt.Errorf("failed to delete all documents: %w", err)
	}
	return nil
}
