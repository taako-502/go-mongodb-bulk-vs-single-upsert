package benchmark

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func UpsertBenchimark(collection *mongo.Collection, count int) (time.Duration, error) {
	if count <= 0 {
		return 0, fmt.Errorf("count must be greater than 0")
	}

	ctx := context.TODO()
	ids, err := seed(ctx, collection, count/2)
	if err != nil {
		return 0, fmt.Errorf("failed to seed data: %w", err)
	}

	// ベンチマーク測定開始
	startTime := time.Now()
	for i, id := range ids {
		upsertData := bson.M{
			"$set": bson.M{
				"text":      "upsert",
				"updatedAt": bson.DateTime(time.Now().UnixNano() / int64(time.Millisecond)),
			},
		}

		// NOTE: upsertは常にtrueに設定
		upsert := options.UpdateOne().SetUpsert(true)
		if i < len(ids)/2 {
			_, err := collection.UpdateOne(
				ctx,
				bson.M{"_id": id}, // 既存のドキュメントをupdateします。
				upsertData,
				upsert,
			)
			if err != nil {
				return 0, fmt.Errorf("failed to update document: %w", err)
			}
		} else {
			_, err := collection.UpdateOne(
				ctx,
				bson.M{"_id": bson.NewObjectID()}, // 新たなドキュメントをinsert
				upsertData,
				upsert,
			)
			if err != nil {
				return 0, fmt.Errorf("failed to insert new document via upsert: %w", err)
			}
		}
	}
	endTime := time.Now()

	return endTime.Sub(startTime), nil
}
