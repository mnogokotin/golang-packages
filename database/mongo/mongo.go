package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Mongo struct {
	Client *mongo.Client
}

func New(connectionUri string, connectTimeout time.Duration) (*Mongo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()
	
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionUri))
	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}

	return &Mongo{
		Client: client,
	}, nil
}

func (m *Mongo) Close() {
	if err := m.Client.Disconnect(context.Background()); err != nil {
		panic(err)
	}
}
