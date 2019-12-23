package database

import (
	"context"

	"github.com/jneubaum/honestvote/tests/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateBlockchain(client *mongo.Client, block Block) bool {
	//Make the block a document and add it to local database
	collection := client.Database("honestvote").Collection(CollectionPrefix + "blockchain")

	document := Block{
		Index:       block.Index,
		Timestamp:   block.Timestamp,
		Transaction: block.Transaction,
		Hash:        block.Hash,
		PrevHash:    block.PrevHash,
		Validator:   block.Validator,
		Valid:       true,
	}

	_, err := collection.InsertOne(context.TODO(), document)

	if err != nil {
		return false
		logger.Println("database_exchange.go", "UpdateBlockchain()", err.Error())
	}

	return true
}