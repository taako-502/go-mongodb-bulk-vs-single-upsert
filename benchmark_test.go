package benchmark_test

import (
	"log"
	"testing"

	"fmt"

	"github.com/taako-502/go-mongodb-bulk-vs-single-upsert/benchmark"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var client *mongo.Client

func init() {
	var err error
	client, err = mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
}

func BenchmarkUpsert(b *testing.B) {
	collection := client.Database("benchmark").Collection("upsert")
	defer benchmark.Cleanup(collection)

	for _, n := range []int{2, 10, 500, 1000, 5000, 10000, 50000, 100000} {
		b.Run("Upsert_"+fmt.Sprint(n), func(b *testing.B) {
			b.ResetTimer()
			for b.Loop() {
				_, err := benchmark.UpsertBenchimark(collection, n)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

func BenchmarkOrderedBulkWrite(b *testing.B) {
	collection := client.Database("benchmark").Collection("ordered_bulk")
	defer benchmark.Cleanup(collection)

	for _, n := range []int{2, 10, 500, 1000, 5000, 10000, 50000, 100000} {
		b.Run("OrderedBulkWrite_"+fmt.Sprint(n), func(b *testing.B) {
			b.ResetTimer()
			for b.Loop() {
				_, err := benchmark.UpsertAndOrderdBulkWriteBenchimark(collection, n)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

func BenchmarkUnorderedBulkWrite(b *testing.B) {
	collection := client.Database("benchmark").Collection("unordered_bulk")
	defer benchmark.Cleanup(collection)

	for _, n := range []int{2, 10, 500, 1000, 5000, 10000, 50000, 100000} {
		b.Run("UnorderedBulkWrite_"+fmt.Sprint(n), func(b *testing.B) {
			b.ResetTimer()
			for b.Loop() {
				_, err := benchmark.UpsertAndUnorderdBulkWriteBenchimark(collection, n)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}
