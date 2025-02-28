package ordered_bulk_write

import (
	"context"
	"fmt"
	"time"

	"github.com/taako-502/go-mongodb-bulk-vs-single-upsert/benchmark"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func InitOrderedBulkWriteModel(ctx context.Context, collection *mongo.Collection, count int) ([]mongo.WriteModel, error) {
	ids, err := benchmark.Seed(ctx, collection, count/2)
	if err != nil {
		return nil, fmt.Errorf("failed to seed data: %w", err)
	}

	var models []mongo.WriteModel
	for i, id := range ids {
		filter := bson.M{"_id": id}
		if i >= len(ids)/2 {
			filter = bson.M{"_id": bson.NewObjectID()}
		}

		update := bson.M{
			"$set": bson.M{
				"text":      "upsert",
				"updatedAt": bson.DateTime(time.Now().UnixNano() / int64(time.Millisecond)),
			},
		}
		model := mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true)
		models = append(models, model)
	}

	return models, nil
}
