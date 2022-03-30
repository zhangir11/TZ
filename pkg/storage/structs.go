package storage

import "go.mongodb.org/mongo-driver/mongo"

type Session struct {
	Guid         string `bson:"guid"`
	RefreshToken string `bson:"refresh_token"`
	ExpiredAt    int64  `bson:"expires_at"`
}

type SessionManager struct {
	collection *mongo.Collection
}

func NewSessionManager(db *mongo.Database, collection string) *SessionManager {
	return &SessionManager{
		collection: db.Collection(collection),
	}
}
