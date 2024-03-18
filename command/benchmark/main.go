package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/taako-502/go-mongodb-bulk-vs-single-upsert/benchmark"
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
		2, 10, 500, 1000, 5000, 10000, 25000, 50000, 60000, 75000, 90000, 100000, 125000, 250000, 500000, 1000000,
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

		// 順序付けられたBulkのベンチマークの実行
		upsertManyCollection := client.Database("benchmark").Collection("orderdBulkWrite")
		bulkWriteDuration, err := benchmark.UpsertAndOrderdBulkWriteBenchimark(upsertManyCollection, p)
		if err != nil {
			panic(err)
		}
		benchmark.Cleanup(upsertManyCollection)

		// 順序付けられていないBulkのベンチマークの実行
		onorderdBulkWriteCollection := client.Database("benchmark").Collection("onorderdBulkWrite")
		unorderdBulkWriteDuration, err := benchmark.UpsertAndUnorderdBulkWriteBenchimark(onorderdBulkWriteCollection, p)
		if err != nil {
			panic(err)
		}
		benchmark.Cleanup(upsertManyCollection)

		// 結果の出力
		durationPrint(p, upsertDuration, bulkWriteDuration, unorderdBulkWriteDuration)
	}
}

func printHeader() {
	fmt.Println("Count,Upsert[ms],Upsert with Orderd BulkWrite[ms],Upsert with Unorderd BulkWrite[ms]")
}

func durationPrint(count int, upsertDuration, bulkWriteDuration, unorderdBulkWriteDuration time.Duration) {
	upsertDurationMs := float64(upsertDuration) / float64(time.Millisecond)
	bulkWriteDurationMs := float64(bulkWriteDuration) / float64(time.Millisecond)
	unorderdBulkWriteDurationMs := float64(unorderdBulkWriteDuration) / float64(time.Millisecond)
	fmt.Printf("%d,%.6f,%.6f,%.6f\n", count, upsertDurationMs, bulkWriteDurationMs, unorderdBulkWriteDurationMs)
}
