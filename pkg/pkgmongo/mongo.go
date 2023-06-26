package pkgmongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongo() (*mongo.Client, error) {
	credential := options.Credential{
		AuthMechanism: "SCRAM-SHA-1",
		Username:      "root",     // your mongodb user
		Password:      "Lumia650", // ...and mongodb
	}
	connString := "mongodb://abc-shard-00-00.gcp.mongodb.net:27017,abc-shard-00-01.gcp.mongodb.net:27017,abc-shard-00-02.gcp.mongodb.net:27017?tls=true"
	clientOpts := options.Client().ApplyURI(connString).SetAuth(credential)
	return mongo.Connect(context.TODO(), clientOpts)
}
