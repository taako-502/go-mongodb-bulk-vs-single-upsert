package benchmark_test

import (
	"context"
	"log"
	"testing"

	"fmt"

	"github.com/taako-502/go-mongodb-bulk-vs-single-upsert/benchmark"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var client *mongo.Client
var benchmarkCounts = []int{2, 10, 500, 1000, 5000, 10000, 50000, 100000}

func init() {
	var err error
	client, err = mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
}

func BenchmarkUpsert(b *testing.B) {
	ctx := context.Background()
	collection := client.Database("benchmark").Collection("upsert")
	defer benchmark.Cleanup(ctx, collection)

	for _, n := range benchmarkCounts {
		b.Run("Upsert_"+fmt.Sprint(n), func(b *testing.B) {
			b.ResetTimer()
			for b.Loop() {
				if _, err := benchmark.UpsertBenchimark(collection, n); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

func BenchmarkOrderedBulkWrite(b *testing.B) {
	ctx := context.Background()
	collection := client.Database("benchmark").Collection("ordered_bulk")
	defer benchmark.Cleanup(ctx, collection)

	for _, n := range benchmarkCounts {
		b.Run("OrderedBulkWrite_"+fmt.Sprint(n), func(b *testing.B) {
			b.ResetTimer()
			model, err := benchmark.InitOrderedBulkWriteModel(ctx, collection, n)
			if err != nil {
				b.Fatal(err)
			}
			for b.Loop() {
				if err := benchmark.UpsertAndOrderdBulkWriteBenchimark(ctx, collection, n, model); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

func BenchmarkUnorderedBulkWrite(b *testing.B) {
	ctx := context.Background()
	collection := client.Database("benchmark").Collection("unordered_bulk")
	defer benchmark.Cleanup(ctx, collection)

	for _, n := range benchmarkCounts {
		b.Run("UnorderedBulkWrite_"+fmt.Sprint(n), func(b *testing.B) {
			b.ResetTimer()
			for b.Loop() {
				if _, err := benchmark.UpsertAndUnorderdBulkWriteBenchimark(collection, n); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}
