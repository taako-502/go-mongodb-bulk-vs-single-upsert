package unordered_bulk_write

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func UpsertAndUnorderedBulkWriteBenchimark(ctx context.Context, collection *mongo.Collection, count int, models []mongo.WriteModel) error {
	// 順序を保証しないことでより高速化させる
	opts := options.BulkWrite().SetOrdered(false)
	if _, err := collection.BulkWrite(ctx, models, opts); err != nil {
		return fmt.Errorf("failed to execute bulk write: %w", err)
	}

	return nil
}
