package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AvailableBlood struct {
	Id         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	BloodGroup string             `bson:"blood_group,omitempty" json:"blood_group,omitempty"`
	Units      int                `bson:"units,omitempty" json:"units,omitempty"`
	Location   string             `bson:"location,omitempty" json:"location,omitempty"`
}

type Donor struct {
	Id           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name         string             `bson:"name,omitempty" json:"name,omitempty"`
	Age          int64              `bson:"age,omitempty" json:"age,omitempty"`
	BloodGroup   string             `bson:"blood_group,omitempty" json:"blood_group,omitempty"`
	Units        int                `bson:"units,omitempty" json:"units,omitempty"`
	DonationTime time.Time          `bson:"donation_time,omitempty" json:"donation_time,omitempty"`
	Location     string             `bson:"location,omitempty" json:"location,omitempty"`
	Adharcard    string             `bson:"adharcard,omitempty" json:"adharcard,omitempty"`
}

type User struct {
	Id         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name       string             `bson:"name,omitempty" json:"name,omitempty"`
	BloodGroup string             `bson:"blood_group,omitempty" json:"blood_group,omitempty"`
	Adharcard  string             `bson:"adharcard,omitempty" json:"adharcard,omitempty"`
	Email      string             `bson:"email,omitempty" json:"email,omitempty"`
	Password   string             `bson:"password,omitempty" json:"password,omitempty"`
	UserId     string             `bson:"user_id,omitempty" json:"user_id,omitempty"`
}

type Patient struct {
	Id             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name           string             `bson:"name,omitempty" json:"name,omitempty"`
	Age            int64              `bson:"age,omitempty" json:"age,omitempty"`
	BloodGroup     string             `bson:"blood_group,omitempty" json:"blood_group,omitempty"`
	Adharcard      string             `bson:"adharcard,omitempty" json:"adharcard,omitempty"`
	Location       string             `bson:"location,omitempty" json:"location,omitempty"`
	RequestedUnits int                `bson:"requested_units,omitempty" json:"requested_units,omitempty"`
	RequestedTime  time.Time          `bson:"requested_time,omitempty" json:"requested_time,omitempty"`
	// RequestClosedTime time.Time          `bson:"reuest_closed_time,omitempty" json:"reuest_closed_time,omitempty"`
	IsBloodProvided bool `bson:"blood_provided,omitempty" json:"blood_provided,omitempty"`
	// Active            bool               `bson:"active,omitempty" json:"active,omitempty"`
}

type Response struct {
	Success    string      `json:"success,omitempty"`
	SucessCode string      `json:"successCode,omitempty"`
	Response   interface{} `json:"response,omitempty"`
}
