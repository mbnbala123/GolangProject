package models

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/membermanagement/store"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Member struct {
	MemberId    uint      `json:"Memberid"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Addresses   Address   `json:"addresses" `
}

type Address struct {
	AddressId uint      `json:"address_id"`
	Street    string    `json:"street"`
	City      string    `json:"city"`
	State     string    `json:"state"`
	Zipcode   string    `json:"postal_code"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var mongoURI string

func init() {
	mongoURI = store.GetMongoDBConnectionString()
}

func MongoConnectionHelper() *mongo.Client {
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (m *Member) InsertMember() {
	client := MongoConnectionHelper()
	defer client.Disconnect(context.TODO())
	collection := client.Database("UMV").Collection("members")
	insertResult, err := collection.InsertOne(context.TODO(), m)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document:", insertResult.InsertedID)
}

func (m *Member) GetAllMembers() []Member {
	var members []Member
	client := MongoConnectionHelper()
	defer client.Disconnect(context.TODO())
	collection := client.Database("UMV").Collection("members")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(context.TODO(), &members); err != nil {
		log.Fatal(err)
	}
	return members
}
func (m *Member) GetMemberByID(memberID uint) *Member {
	var member Member
	client := MongoConnectionHelper()
	defer client.Disconnect(context.TODO())
	collection := client.Database("UMV").Collection("members")
	filter := bson.D{{Key: "memberid", Value: memberID}}
	err := collection.FindOne(context.TODO(), filter).Decode(&member)
	if err != nil {
		log.Fatal(err)
	}
	return &member
}
func (m *Member) UpdateMember(memberID uint) {
	client := MongoConnectionHelper()
	defer client.Disconnect(context.TODO())
	collection := client.Database("UMV").Collection("members")
	filter := bson.D{{Key: "memberid", Value: memberID}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "first_name", Value: m.FirstName},
			{Key: "last_name", Value: m.LastName},
			{Key: "email", Value: m.Email},
			{Key: "phone_number", Value: m.PhoneNumber},
			{Key: "updated_at", Value: time.Now()},
			{Key: "addresses", Value: m.Addresses},
		}},
	}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Matched and modified documents:", updateResult.MatchedCount, updateResult.ModifiedCount)
}
func (m *Member) DeleteMember(memberID uint) {
	client := MongoConnectionHelper()
	defer client.Disconnect(context.TODO())
	collection := client.Database("UMV").Collection("members")
	filter := bson.D{{Key: "memberid", Value: memberID}}
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted documents:", deleteResult.DeletedCount)
}
