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

// seed ベンチマーク測定に利用するデータ作成
func seed(ctx context.Context, collection *mongo.Collection, numRecords int) ([]primitive.ObjectID, error) {
	var models []mongo.WriteModel
	var ids []primitive.ObjectID

	for range numRecords {
		id := primitive.NewObjectID()
		ids = append(ids, id)
		model := mongo.NewInsertOneModel().SetDocument(bson.M{
			"_id":       id,
			"text":      "initial",
			"createdAt": primitive.DateTime(time.Now().UnixNano() / int64(time.Millisecond)),
		})
		models = append(models, model)
	}

	// 一括操作の実行
	opts := options.BulkWrite().SetOrdered(false) // 順序を保証しないことでより高速化させる
	if _, err := collection.BulkWrite(ctx, models, opts); err != nil {
		return nil, fmt.Errorf("failed to perform bulk insert: %w", err)
	}

	return ids, nil
}
