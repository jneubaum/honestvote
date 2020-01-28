package database

import (
	"context"
	"strconv"

	"github.com/jneubaum/honestvote/tests/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddBlock(block Block) error {
	//Make the block a document and add it to local database
	collection := MongoDB.Database("honestvote").Collection(CollectionPrefix + "blockchain")

	_, err := collection.InsertOne(context.TODO(), block)

	if err != nil {
		logger.Println("database_exchange.go", "UpdateBlockchain()", err.Error())
		return err
	}

	return nil
}

func LastIndex(client *mongo.Client) int64 {
	collection := client.Database("honestvote").Collection(CollectionPrefix + "blockchain")

	index, err := collection.CountDocuments(context.TODO(), bson.M{})

	if err == nil {
		return index
	}

	return 0
}

func GrabDocuments(client *mongo.Client, old_index string) []Block {

	var blocks []Block

	collection := client.Database("honestvote").Collection(CollectionPrefix + "blockchain")

	index, _ := strconv.Atoi(old_index)
	current, err := collection.CountDocuments(context.TODO(), bson.M{})

	if current-int64(index) > 0 && err == nil {
		result, err := collection.Find(context.TODO(), bson.M{"index": bson.M{"$gt": index}})

		if err != nil {
			logger.Println("database_exchange.go", "GrabDocuments()", err.Error())
		}

		for result.Next(context.TODO()) {
			var block Block
			err = result.Decode(&block)
			blocks = append(blocks, block)
		}

		return blocks
	}

	return nil
}

func UpdateMongo(client *mongo.Client, data []Block) {
	collection := client.Database("honestvote").Collection(CollectionPrefix + "blockchain")

	var ui []interface{}
	for _, block := range data {
		ui = append(ui, block)
	}

	collection.InsertMany(context.TODO(), ui)
}
