package ordered_bulk_write

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func UpsertAndOrderdBulkWriteBenchimark(ctx context.Context, collection *mongo.Collection, count int, models []mongo.WriteModel) error {
	opts := options.BulkWrite().SetOrdered(true)
	if _, err := collection.BulkWrite(ctx, models, opts); err != nil {
		return fmt.Errorf("failed to execute bulk write: %w", err)
	}

	return nil
}
