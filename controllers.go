package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
)

var db = dbClient()
var singleplayerCollection = db.Database("Stats").Collection("singleplayer")
var multiplayerCollection = db.Database("Stats").Collection("multiplayer")

type Response struct {
	Status string
	Data   interface{}
}

func okResponse(w http.ResponseWriter, data interface{}, status int) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(Response{Status: "OK", Data: data})
}

func errResponse(w http.ResponseWriter, data interface{}, status int) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(Response{Status: "ERROR", Data: data})
}

func InsertSingleplayer(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var asJson interface{}
	err := decoder.Decode(&asJson)
	if err != nil {
		errResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	insertRes, err := singleplayerCollection.InsertOne(context.TODO(), asJson)
	if err != nil {
		errResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("Inserted single-player: %v+", insertRes)
	// TODO better alternative to sprintf'ing?
	okResponse(w, fmt.Sprintf("%v", insertRes), http.StatusCreated)

}

func InsertMultiplayer(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var asJson interface{}
	err := decoder.Decode(&asJson)
	if err != nil {
		errResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	insertRes, err := multiplayerCollection.InsertOne(context.TODO(), asJson)
	if err != nil {
		errResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("Inserted multiplayer: %v+", insertRes)
	okResponse(w, fmt.Sprintf("%v", insertRes), http.StatusCreated)
}

func GetMultiplayerLeaderboard(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	sortBy := r.FormValue("sortBy")
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
		errResponse(w, err.Error(), http.StatusInternalServerError)
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
		errResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	okResponse(w, results, http.StatusOK)
}

func GetSingleplayerLeaderboard(w http.ResponseWriter, r *http.Request) {
	ctx := context.TODO()
	sortBy := r.FormValue("sortBy")
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
		errResponse(w, err.Error(), http.StatusInternalServerError)
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
		errResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	okResponse(w, results, http.StatusOK)
}
