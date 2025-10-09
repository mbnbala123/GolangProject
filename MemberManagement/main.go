package main

import (
	"fmt"
	"time"

	"github.com/membermanagement/models"
	"github.com/membermanagement/store"
)

func GetCurrentTime() time.Time {
	return time.Now()
}
func main() {
	fmt.Println("Hello, World!")
	connectionString := store.GetMongoDBConnectionString()
	fmt.Println("MongoDB Connection String:", connectionString)

	member := models.Member{
		MemberId:    1,
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "john.doe@example.com",
		PhoneNumber: "123-456-7890",
		CreatedAt:   time.Now(),
		Addresses: models.Address{
			AddressId: 1,
			Street:    "123 Main St",
			City:      "Anytown",
			State:     "CA",
			Zipcode:   "12345",
			Country:   "USA",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	member.InsertMember()
	members := member.GetAllMembers()
	fmt.Println("Members in the database:", members)
}
