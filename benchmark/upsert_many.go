package benchmark

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpsertManyBenchimark(collection *mongo.Collection, count int) (time.Duration, error) {
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

	// データの作成
	var models []mongo.WriteModel
	for i, id := range ids {
		filter := bson.M{"_id": id}
		if i >= len(ids)/2 {
			// 新たなドキュメントを挿入するため、IDを新しく生成します。
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
	startTime := time.Now()
	// BulkWrite操作を実行します。
	_, err := collection.BulkWrite(ctx, models)
	if err != nil {
		return 0, fmt.Errorf("failed to execute bulk write: %w", err)
	}
	endTime := time.Now()

	return endTime.Sub(startTime), nil
}
