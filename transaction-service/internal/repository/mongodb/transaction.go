package mongodb

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/walbety/transaction-app/transaction-service/internal/canonical"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	database = "transaction"
	collection = "purchase"

)

func (d MongoDBImpl) SaveTransaction(ctx context.Context, transaction *canonical.Transaction) (string,error) {

	coll := d.client.Database(database).Collection(collection)
	result, err := coll.InsertOne(ctx, transaction)
	if err != nil {
		log.WithFields(log.Fields{
			"transaction.Date":transaction.Date,
			"transaction.Description":transaction.Description,
		}).WithError(err).Error("Error at insertion - repository layer")
		return "",err
	}

	return result.InsertedID.(primitive.ObjectID).String(), nil
}

func (d MongoDBImpl) FindTransactionById(ctx context.Context, id string) (canonical.Transaction, error) {

	coll := d.client.Database(database).Collection(collection)

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		log.Println("Invalid id")
		return canonical.Transaction{}, err
	}

	result:= coll.FindOne(ctx, bson.M{"_id": objectId})
	transaction := canonical.Transaction{}
	err = result.Decode(transaction)

	if err != nil {
		log.WithContext(ctx).
			WithError(err).
			WithField("id", id).
			Error("failed to retrieve transaction")
		return canonical.Transaction{}, err
	}
	return transaction, nil
}

