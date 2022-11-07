package model

type Address struct {
	Street  string `bson:"street,omitempty" json:"street,omitempty"`
	City    string `bson:"city,omitempty" json:"city,omitempty"`
	State   string `bson:"state,omitempty" json:"state,omitempty"`
	Zipcode string `bson:"zipcode,omitempty" json:"zipcode,omitempty"`
	Country string `bson:"country,omitempty" json:"country,omitempty"`
}
