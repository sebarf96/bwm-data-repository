package user_repository

import (
	"bwm-api-repository/src/database"
	"bwm-api-repository/src/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"context"

	"go.mongodb.org/mongo-driver/bson"
)

var (
	collection = database.GetCollection("users")
	ctx        = context.Background()
)

type UserRepository interface {
	Create(user models.User) error
	Read() (models.Users, error)
	Update(user models.User, id string) error
	Delete(id string) error
}

type userRepository struct {
}

func NewUserRepository(*mongo.Collection) UserRepository {
	return &userRepository{}
}

func (u *userRepository) Create(user models.User) error {

	var err error
	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) Read() (models.Users, error) {

	var users models.Users
	filter := bson.D{}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {

		var user models.User
		err = cur.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)

	}
	return users, nil
}

func (u *userRepository) Update(user models.User, id string) error {

	oid, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": oid}
	update := bson.M{
		"$set": bson.M{
			"name":      user.Name,
			"email":     user.Email,
			"update_at": time.Now(),
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepository) Delete(id string) error {

	oid, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": oid}
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
