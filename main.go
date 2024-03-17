package main

import (
	"context"
	"fmt"
	"time"

	"github.com/taako-502/go-mongodb-upsert-vs-upsertmany/benchmark"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	fmt.Println("--------------------------------------------------------------------")
	fmt.Println("")

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	count := 1000
	fmt.Println("処理回数\t: ", count)

	// upsertのベンチマークの実行
	upsertCollection := client.Database("benchmark").Collection("upsert")
	duration, err := benchmark.UpsertBenchimark(upsertCollection, count)
	if err != nil {
		panic(err)
	}
	durationPrint("upsert", duration)
	benchmark.Cleanup(upsertCollection)

	// upsertManyのベンチマークの実行
	upsertManyCollection := client.Database("benchmark").Collection("upsertMany")
	duration, err = benchmark.UpsertAndBulkWriteBenchimark(upsertManyCollection, count)
	if err != nil {
		panic(err)
	}
	benchmark.Cleanup(upsertManyCollection)
	durationPrint("upsert many", duration)

	fmt.Println("")
	fmt.Println("--------------------------------------------------------------------")
}

func durationPrint(target string, duration time.Duration) {
	println(target + "の処理時間\t: " + duration.String())
}
