package benchmark

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpsertAndOrderdBulkWriteBenchimark(collection *mongo.Collection, count int) (time.Duration, error) {
	if count <= 0 {
		return 0, fmt.Errorf("count must be greater than 0")
	}

	ctx := context.TODO()
	ids, err := seed(ctx, collection, count/2)
	if err != nil {
		return 0, fmt.Errorf("failed to seed data: %w", err)
	}

	// データの作成
	var models []mongo.WriteModel
	for i, id := range ids {
		filter := bson.M{"_id": id}
		if i >= len(ids)/2 {
			filter = bson.M{"_id": primitive.NewObjectID()}
		}

		update := bson.M{
			"$set": bson.M{
				"text":      "upsert",
				"updatedAt": primitive.DateTime(time.Now().UnixNano() / int64(time.Millisecond)),
			},
		}
		model := mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true)
		models = append(models, model)
	}

	// ベンチマーク測定開始
	opts := options.BulkWrite().SetOrdered(true) // 順序を保証する
	startTime := time.Now()
	if _, err := collection.BulkWrite(ctx, models, opts); err != nil {
		return 0, fmt.Errorf("failed to execute bulk write: %w", err)
	}
	endTime := time.Now()

	return endTime.Sub(startTime), nil
}
