package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
)

var db = dbClient()
var singleplayerCollection = db.Database("Stats").Collection("singleplayer")
var multiplayerCollection = db.Database("Stats").Collection("multiplayer")

func insertSingleplayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	insertRes, err := singleplayerCollection.InsertOne(context.TODO(), r.Body)
	if err != nil {
		fmt.Print(err)
		return
		// Failurre status and message
	}
	log.Printf("Inserted single-player: %v+", insertRes)
}

func insertMultiplayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	insertRes, err := multiplayerCollection.InsertOne(context.TODO(), r.Body)
	if err != nil {
		fmt.Print(err)
		return
	}
	log.Printf("Inserted multiplayer: %v+", insertRes)
}

func getMultiplayerLeaderboard(w http.ResponseWriter, r *http.Request) {
	return
}

func getSingleplayerLeaderboard(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	w.Header().Set("Content-Type", "application/json")
	sortBy := "$highestScore"
	group := bson.M{
		"$group": bson.M{
			"_id":          "$userID",
			"highestScore": bson.M{"$max": "$score"},
			"longestLife":  bson.M{"$max": "$lifetime"},
		},
	}
	sort := bson.M{
		"$sort": bson.M{sortBy: -1},
	}
	cur, err := singleplayerCollection.Aggregate(ctx, []bson.M{
		group,
		sort,
	})
	if err != nil {
		fmt.Print(err)
		return
	}
	defer cur.Close(ctx)
	/*cursor.Next(ctx)
	elem := &bson.D{}
	err = cursor.Decode(elem)
	m := elem.Map()
	results := m["allkeys"].(primitive.A)*/
	results := []bson.M{}
	if err := cur.All(ctx, &results); err != nil {
		fmt.Print(err)
		return
	}
}
