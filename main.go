package main

import (
	"context"
	"fmt"
	"time"

	"github.com/taako-502/go-mongodb-upsert-vs-upsertmany/upsert"
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
	startTime := time.Now()
	if err := upsert.Upsert(upsertCollection, count); err != nil {
		panic(err)
	}
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	durationPrint("upsert", duration)

	// upsertManyのベンチマークの実行
	upsertManyCollection := client.Database("benchmark").Collection("upsertMany")
	startTime = time.Now()
	if err := upsert.UpsertMany(upsertManyCollection, count); err != nil {
		panic(err)
	}
	endTime = time.Now()
	duration = endTime.Sub(startTime)
	durationPrint("upsert many", duration)

	fmt.Println("")
	fmt.Println("--------------------------------------------------------------------")
}

func durationPrint(target string, duration time.Duration) {
	println(target + "の処理時間\t: " + duration.String())
}
