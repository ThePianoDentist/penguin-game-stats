package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func dbClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	return client
}

type MultiplayerResult struct {
	CreatedAt      uint64
	FinishedAt     uint64
	Type           string
	Cancelled      bool
	TeamOneScore   uint32
	TeamTwoScore   uint32
	TeamOnePlayers []PlayerResult
	TeamTwoPlayers []PlayerResult
}

type PlayerResult struct {
	PlayerID              string
	Win                   int32 // 0 for cancelled. +1 for win. -1 for loss
	Kills                 uint32
	Deaths                uint32
	DaggerKills           uint32
	DaggerReflectionKills uint32
	ShockwaveKills        uint32
	RightClickKills       uint32
}
