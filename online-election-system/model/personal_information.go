package model

type PersonalInfo struct {
	Name         string  `bson:"name,omitempty" json:"name,omitempty"`
	FatherName   string  `bson:"father_name,omitempty" json:"father_name,omitempty"`
	DOB          string  `bson:"dob,omitempty" json:"dob,omitempty"`
	Age          int     `bson:"age,omitempty" json:"age,omitempty"`
	VoterId      string  `bson:"voter_id,omitempty" json:"voter_id,omitempty"`
	DocumentType string  `bson:"document_type,omitempty" json:"document_type,omitempty"`
	Address      Address `bson:"address,omitempty" json:"address,omitempty"`
}
