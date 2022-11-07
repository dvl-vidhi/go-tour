package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Candidates struct {
	UserId           primitive.ObjectID `bson:"user_id" json:"user_id"`
	Name             string             `bson:"name" json:"name"`
	Commitments      []string           `bson:"commitments" json:"commitments"`
	Votecount        string             `bson:"vote_count" json:"vote_count"`
	VoteSign         string             `bson:"vote_sign" json:"vote_sign"`
	NominationStatus string             `bson:"nomination_status" json:"nomination_status"`
	VerifiedBy       VerifiedBy         `bson:"verified_by,omitempty" json:"verified_by,omitempty"`
}

type CandidatesRequest struct {
	ElectionId       string   `json:"election_id,omitempty"`
	UserId           string   `json:"user_id,omitempty"`
	Name             string   `json:"name,omitempty"`
	Commitments      []string `json:"commitments,omitempty"`
	VoteSign         string   `json:"vote_sign,omitempty"`
	NominationStatus string   `bson:"nomination_status" json:"nomination_status"`
	Email            string   `bson:"email" json:"email"`
}
