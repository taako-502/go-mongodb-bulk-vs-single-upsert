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

func UpsertBenchimark(collection *mongo.Collection, count int) (time.Duration, error) {
	ctx := context.TODO()

	var ids []primitive.ObjectID
	for i := 0; i < count/2; i++ {
		ids = append(ids, primitive.NewObjectID())
	}

	// テストデータ作成
	for _, id := range ids {
		_, err := collection.InsertOne(ctx, bson.M{
			"_id":       id,
			"text":      "initial",
			"createdAt": primitive.DateTime(time.Now().UnixNano() / int64(time.Millisecond)),
		})
		if err != nil {
			return 0, fmt.Errorf("failed to insert initial document: %w", err)
		}
	}

	// ベンチマーク測定開始
	startTime := time.Now()
	for i, id := range ids {
		upsertData := bson.M{
			"$set": bson.M{
				"text":      "upsert",
				"updatedAt": primitive.DateTime(time.Now().UnixNano() / int64(time.Millisecond)),
			},
		}

		// NOTE: upsertは常にtrueに設定
		upsert := options.Update().SetUpsert(true)
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
				bson.M{"_id": primitive.NewObjectID()}, // 新たなドキュメントをinsert
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
