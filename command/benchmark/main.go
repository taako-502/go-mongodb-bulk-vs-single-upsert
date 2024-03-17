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

type pattern struct {
	numberOfUsers  int
	numberOfPoints int
}

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
		10, 100, 1000, 10000, 100000, 1000000,
	}

	for _, p := range patterns {
		fmt.Println("処理回数\t: ", p)

		// upsertのベンチマークの実行
		upsertCollection := client.Database("benchmark").Collection("upsert")
		duration, err := benchmark.UpsertBenchimark(upsertCollection, p)
		if err != nil {
			panic(err)
		}
		durationPrint("upsert", duration)
		benchmark.Cleanup(upsertCollection)

		// upsertManyのベンチマークの実行
		upsertManyCollection := client.Database("benchmark").Collection("upsertMany")
		duration, err = benchmark.UpsertAndBulkWriteBenchimark(upsertManyCollection, p)
		if err != nil {
			panic(err)
		}
		durationPrint("upsert many", duration)
		benchmark.Cleanup(upsertManyCollection)
	}
}

func durationPrint(target string, duration time.Duration) {
	println(target + "の処理時間\t: " + duration.String())
}
