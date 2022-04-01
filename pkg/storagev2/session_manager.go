package storage

import (
	"authentication/config"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//-------------------------------------------------------------------------------------------------------
func SetupConnection() *mongo.Database {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(config.Conf.DatabaseURI).
		SetServerAPIOptions(serverAPIOptions)
	log.Println(config.Conf.DatabaseURI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println("Mango Connection Error:")
		log.Fatal(err)
	}

	if err := client.Connect(ctx); err != nil && err.Error() != "topology is connected or connecting" {
		log.Fatal("Mango Client Connection Error: " + err.Error())
	}

	if err = client.Ping(context.Background(), nil); err != nil {
		log.Fatal("Ping is very high: " + err.Error())
	}

	return client.Database(config.Conf.DatabaseName)
}

//-------------------------------------------------------------------------------------------------------

func (s *SessionManager) Insert(value *Session) error {
	if _, err := s.collection.InsertOne(context.Background(), value); err != nil {
		return err
	}

	return nil
}

//-------------------------------------------------------------------------------------------------------

func (s *SessionManager) Get(guid string) (*Session, error) {
	var result Session

	if err := s.collection.FindOne(context.TODO(), bson.D{{"guid", guid}}).Decode(&result); err != nil {
		return &Session{}, err
	}

	return &result, nil
}

//-------------------------------------------------------------------------------------------------------

func (s *SessionManager) Delete(guid string) error {
	if _, err := s.collection.DeleteOne(context.Background(), bson.D{{"guid", guid}}); err != nil {
		return err
	}

	return nil
}
