package benchmark

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// seed ベンチマーク測定に利用するデータ作成
func seed(ctx context.Context, collection *mongo.Collection, numRecords int) ([]primitive.ObjectID, error) {
	var ids []primitive.ObjectID
	for range numRecords {
		ids = append(ids, primitive.NewObjectID())
	}

	for _, id := range ids {
		_, err := collection.InsertOne(ctx, bson.M{
			"_id":       id,
			"text":      "initial",
			"createdAt": primitive.DateTime(time.Now().UnixNano() / int64(time.Millisecond)),
		})
		if err != nil {
			return nil, fmt.Errorf("failed to insert initial document: %w", err)
		}
	}

	return ids, nil
}
