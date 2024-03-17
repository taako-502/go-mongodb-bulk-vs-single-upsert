package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/taako-502/go-mongodb-upsert-vs-upsertmany/benchmark"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	// 計測パターン
	patterns := []int{
		2, 10, 100, 1000, 10000, 100000, 1000000,
	}

	printHeader()
	for _, p := range patterns {
		// upsertのベンチマークの実行
		upsertCollection := client.Database("benchmark").Collection("upsert")
		upsertDuration, err := benchmark.UpsertBenchimark(upsertCollection, p)
		if err != nil {
			panic(err)
		}
		benchmark.Cleanup(upsertCollection)

		// upsertManyのベンチマークの実行
		upsertManyCollection := client.Database("benchmark").Collection("upsertMany")
		bulkWriteDuration, err := benchmark.UpsertAndBulkWriteBenchimark(upsertManyCollection, p)
		if err != nil {
			panic(err)
		}
		durationPrint(p, upsertDuration, bulkWriteDuration)
		benchmark.Cleanup(upsertManyCollection)
	}
}

func printHeader() {
	fmt.Println("Count\tUpsert[ms]\tUpsert with BulkWrite[ms]")
}

func durationPrint(count int, upsertDuration, bulkWriteDuration time.Duration) {
	upsertDurationMs := float64(upsertDuration) / float64(time.Millisecond)
	bulkWriteDurationMs := float64(bulkWriteDuration) / float64(time.Millisecond)
	fmt.Printf("%d\t%.6f\t%.6f\n", count, upsertDurationMs, bulkWriteDurationMs)
}
