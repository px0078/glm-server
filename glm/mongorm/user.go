package mongorm

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Model struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

type User struct {
	Model
	Username string `bson:"username"`
	Password string `bson:"password"`
	Token    string `bson:"token"`
}
