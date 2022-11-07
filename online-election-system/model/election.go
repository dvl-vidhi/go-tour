package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Election struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Location       string             `bson:"location,omitempty" json:"location,omitempty"`
	ElectionDate   string             `bson:"election_date,omitempty" json:"election_date,omitempty"`
	ResultDate     string             `bson:"result_date,omitempty" json:"result_date,omitempty"`
	Result         string             `bson:"result,omitempty" json:"result,omitempty"`
	ElectionStatus string             `bson:"election_status,omitempty" json:"election_status,omitempty"`
	Candidates     []Candidates       `bson:"candidates,omitempty" json:"candidates,omitempty"`
}
