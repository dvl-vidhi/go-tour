package dao

import (
	"context"
	"errors"
	"fmt"
	"log"

	"online-election-system/helper"
	"online-election-system/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ElectionDAO struct {
	Server     string
	Database   string
	Collection string
}

var ElectionCollection *mongo.Collection
var Electionctx = context.TODO()

func (e *ElectionDAO) ElectionConnect() {
	clientOptions := options.Client().ApplyURI(e.Server)
	client, err := mongo.Connect(Electionctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(Electionctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	ElectionCollection = client.Database(e.Database).Collection(e.Collection)
}

func (e *ElectionDAO) ElectionInsert(Election model.Election) (model.Election, error) {
	data, err := ElectionCollection.InsertOne(Electionctx, Election)

	if err != nil {
		fmt.Print(err.Error())
		return Election, errors.New("unable to create new record")
	}

	if oid, ok := data.InsertedID.(primitive.ObjectID); ok {

		Election.ID = oid

	}

	return Election, nil
}

func (e *ElectionDAO) AddCandidate(req model.CandidatesRequest) (bson.M, error) {
	var UpdatedElection bson.M
	var Elections []*model.Election
	electionId, err := primitive.ObjectIDFromHex(req.ElectionId)
	if err != nil {
		return UpdatedElection, err
	}
	cur, err := ElectionCollection.Find(Electionctx, bson.D{primitive.E{Key: "_id", Value: electionId}})

	if err != nil {
		return UpdatedElection, errors.New("unable to query db")
	}
	UserIdHex, err := primitive.ObjectIDFromHex(req.UserId)
	if err != nil {
		return UpdatedElection, err
	}
	_, err = UserCollection.Find(Userctx, bson.D{primitive.E{Key: "_id", Value: UserIdHex}})
	if err != nil {
		return UpdatedElection, err
	}
	var uploadPath = "upload/voteSign/"
	msg, err := helper.UploadFile(req.VoteSign, uploadPath)
	if err != nil {
		log.Println(err)
		return UpdatedElection, errors.New("Unable to upload file")
	}
	fmt.Println("Upload file:", msg)

	Elections, err = convertDbResultIntoElectionStruct(cur)
	if err != nil {
		return UpdatedElection, err
	}

	if len(Elections) == 0 {
		return UpdatedElection, mongo.ErrNoDocuments
	}

	var Candidates = Elections[0].Candidates

	for i := range Candidates {
		fmt.Println(Candidates[i].UserId.String())
		if Candidates[i].UserId == UserIdHex {
			return UpdatedElection, errors.New("Candidate already registered for the given election")
		}
	}

	filter := bson.D{primitive.E{Key: "_id", Value: electionId}}
	UpdateQuery := bson.D{}
	UpdateQuery = append(UpdateQuery, primitive.E{Key: "user_id", Value: UserIdHex})
	UpdateQuery = append(UpdateQuery, primitive.E{Key: "name", Value: req.Name})
	UpdateQuery = append(UpdateQuery, primitive.E{Key: "commitments", Value: req.Commitments})
	UpdateQuery = append(UpdateQuery, primitive.E{Key: "vote_sign", Value: req.VoteSign})
	UpdateQuery = append(UpdateQuery, primitive.E{Key: "nomination_status", Value: "not verified"})

	update := bson.D{{Key: "candidates", Value: UpdateQuery}}
	update = bson.D{{Key: "$push", Value: update}}

	ElectionCollection.FindOneAndUpdate(Electionctx, filter, update).Decode(&UpdatedElection)

	_, err = ElectionCollection.Find(Electionctx, filter)
	if err != nil {
		return UpdatedElection, err
	}

	return UpdatedElection, nil
}

func (e *ElectionDAO) ElectionFindById(id string) ([]*model.Election, error) {
	var Elections []*model.Election

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Elections, err
	}
	cur, err := ElectionCollection.Find(Electionctx, bson.D{primitive.E{Key: "_id", Value: idHex}})

	if err != nil {
		return Elections, errors.New("unable to query db")
	}

	Elections, err = convertDbResultIntoElectionStruct(cur)

	if err != nil {
		return Elections, err
	}

	if len(Elections) == 0 {
		return Elections, mongo.ErrNoDocuments
	}

	return Elections, nil
}

func (e *ElectionDAO) FilterOnElectionDetails(req model.Election) ([]*model.Election, error) {
	var Elections []*model.Election
	query := bson.D{}

	if req.Location != "" {
		query = append(query, primitive.E{Key: "location", Value: req.Location})
	}
	if req.Result != "" {
		query = append(query, primitive.E{Key: "result", Value: req.Result})
	}
	if req.ElectionStatus != "" {
		query = append(query, primitive.E{Key: "election_status", Value: req.ElectionStatus})
	}
	if req.ElectionDate != "" {
		query = append(query, primitive.E{Key: "election_date", Value: req.ElectionDate})
	}
	if req.ResultDate != "" {
		query = append(query, primitive.E{Key: "result_date", Value: req.ElectionDate})
	}

	cur, err := ElectionCollection.Find(Electionctx, query)

	if err != nil {
		return Elections, errors.New("unable to query db")
	}

	Elections, err = convertDbResultIntoElectionStruct(cur)

	if err != nil {
		return Elections, err
	}

	if len(Elections) == 0 {
		return Elections, mongo.ErrNoDocuments
	}

	return Elections, nil

}

func (e *ElectionDAO) Update(election model.Election, id string) (bson.M, error) {
	var updatedElection bson.M
	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return updatedElection, err
	}

	filter := bson.D{primitive.E{Key: "_id", Value: idHex}}

	update := bson.D{primitive.E{Key: "$set", Value: election}}

	err = ElectionCollection.FindOneAndUpdate(Electionctx, filter, update).Decode(&updatedElection)
	if err != nil {
		return updatedElection, err
	}

	if updatedElection == nil {
		return updatedElection, errors.New("Data not present in db given by Id or it is deactivated")
	}

	return updatedElection, nil
}

func (e *ElectionDAO) VerifyCandidate(candidatesRequest model.CandidatesRequest, id string) (bson.M, error) {
	var updatedElection bson.M
	idHex, err := primitive.ObjectIDFromHex(id)
	var adminData []*model.User

	data, err := UserCollection.Find(Userctx, bson.D{primitive.E{Key: "email", Value: candidatesRequest.Email}})
	adminData, err = convertDbResultIntoUserStruct(data)
	if err != nil {
		return updatedElection, err
	}
	filter := bson.D{}

	var Elections []*model.Election

	cur, err := ElectionCollection.Find(Electionctx, bson.D{primitive.E{Key: "_id", Value: idHex}})

	if err != nil {
		return updatedElection, errors.New("unable to query db")
	}

	Elections, err = convertDbResultIntoElectionStruct(cur)

	if err != nil {
		return updatedElection, err
	}

	if len(Elections) == 0 {
		return updatedElection, mongo.ErrNoDocuments
	}

	var Candidates = Elections[0].Candidates

	UserIdHex, err := primitive.ObjectIDFromHex(candidatesRequest.UserId)

	for i := range Candidates {
		fmt.Println(Candidates[i].UserId.String())
		if Candidates[i].UserId == UserIdHex {
			Candidates[i].NominationStatus = candidatesRequest.NominationStatus
			Candidates[i].VerifiedBy.Name = adminData[0].Name
			Candidates[i].VerifiedBy.Id = adminData[0].ID
		}
	}

	filter = bson.D{primitive.E{Key: "_id", Value: idHex}}
	update := bson.D{primitive.E{Key: "$set", Value: Elections[0]}}

	err = ElectionCollection.FindOneAndUpdate(Electionctx, filter, update).Decode(&updatedElection)
	if err != nil {
		return updatedElection, err
	}
	fmt.Println(updatedElection)
	if updatedElection == nil {
		return updatedElection, errors.New("Data not present in db given by Id or it is deactivated")
	}
	return updatedElection, err
}

func convertDbResultIntoElectionStruct(fetchDataCursor *mongo.Cursor) ([]*model.Election, error) {
	var finaldata []*model.Election
	for fetchDataCursor.Next(Userctx) {
		var data model.Election
		err := fetchDataCursor.Decode(&data)
		if err != nil {
			return finaldata, err
		}
		finaldata = append(finaldata, &data)
	}
	return finaldata, nil
}
