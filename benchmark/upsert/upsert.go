package upsert

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func UpsertBenchmark(ctx context.Context, collection *mongo.Collection, ids []bson.ObjectID) error {
	// ベンチマーク測定開始
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
				return fmt.Errorf("failed to update document: %w", err)
			}
		} else {
			_, err := collection.UpdateOne(
				ctx,
				bson.M{"_id": bson.NewObjectID()}, // 新たなドキュメントをinsert
				upsertData,
				upsert,
			)
			if err != nil {
				return fmt.Errorf("failed to insert new document via upsert: %w", err)
			}
		}
	}

	return nil
}
