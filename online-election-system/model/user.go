package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Role         string               `bson:"role" json:"role"`
	Name         string               `bson:"name" json:"name"`
	Email        string               `bson:"email" json:"email"`
	Password     string               `bson:"password" json:"-password"`
	PhoneNumber  string               `bson:"phone_number" json:"phone_number"`
	IsVerified   bool                 `bson:"is_verified" json:"is_verified"`
	VerifiedBy   VerifiedBy           `bson:"verified_by,omitempty" json:"verified_by,omitempty"`
	UploadedDocs UploadedDocs         `bson:"uploaded_docs,omitempty" json:"uploaded_docs,omitempty"`
	PersonalInfo PersonalInfo         `bson:"personal_info,omitempty" json:"personal_info,omitempty"`
	Voted        []primitive.ObjectID `bson:"voted" json:"voted"`
}
