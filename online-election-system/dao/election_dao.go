package dao

import (
	"context"
	"errors"
	"fmt"
	"log"

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
	userId, err := primitive.ObjectIDFromHex(req.UserId)
	if err != nil {
		return UpdatedElection, err
	}
	_, err = UserCollection.Find(Userctx, bson.D{primitive.E{Key: "_id", Value: userId}})
	if err != nil {
		return UpdatedElection, err
	}

	msg, err := UploadFile(req.VoteSign)
	if err != nil {
		log.Println(err)
		return UpdatedElection, errors.New("Unable to upload file")
	}
	fmt.Println("Upload file:", msg)

	for cur.Next(Electionctx) {
		var e model.Election

		err := cur.Decode(&e)

		if err != nil {
			return UpdatedElection, err
		}

		Elections = append(Elections, &e)
	}
	var Candidates []model.Candidates
	Candidates = Elections[0].Candidates
	// Candidates[].find
	UserIdHex, err := primitive.ObjectIDFromHex(req.UserId)

	for i := range Candidates {
		fmt.Println(Candidates[i].UserId.String())
		if Candidates[i].UserId == UserIdHex {
			return UpdatedElection, errors.New("Candidate already registered for the given election")
		}
	}

	filter := bson.D{primitive.E{Key: "_id", Value: electionId}}
	UpdateQuery := bson.D{}
	UpdateQuery = append(UpdateQuery, primitive.E{Key: "user_id", Value: userId})
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

	for cur.Next(Electionctx) {
		var e model.Election

		err := cur.Decode(&e)

		if err != nil {
			return Elections, err
		}

		Elections = append(Elections, &e)
	}

	if err := cur.Err(); err != nil {
		return Elections, err
	}

	cur.Close(Electionctx)

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
		// electionDate, err := convertDate(req.ElectionDate)
		// if err != nil {
		// 	return Elections, "Error Occurred", err
		// }
		query = append(query, primitive.E{Key: "election_date", Value: req.ElectionDate})
	}
	if req.ResultDate != "" {
		// resultDate, err := convertDate(req.ResultDate)
		// if err != nil {
		// 	return Elections, "Error Occurred", err
		// }
		query = append(query, primitive.E{Key: "result_date", Value: req.ElectionDate})
	}

	cur, err := ElectionCollection.Find(Electionctx, query)

	if err != nil {
		return Elections, errors.New("unable to query db")
	}

	for cur.Next(Electionctx) {
		var e model.Election

		err := cur.Decode(&e)

		if err != nil {
			return Elections, err
		}

		Elections = append(Elections, &e)
	}

	if err := cur.Err(); err != nil {
		return Elections, err
	}

	cur.Close(Electionctx)

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

	// e := &model.User{}
	err = ElectionCollection.FindOneAndUpdate(Electionctx, filter, update).Decode(&updatedElection)
	if err != nil {
		return updatedElection, err
	}

	if updatedElection == nil {
		return updatedElection, errors.New("Data not present in db given by Id or it is deactivated")
	}

	return updatedElection, nil
}

func (epd *ElectionDAO) VerifyCandidate(candidatesRequest model.CandidatesRequest, id string) (bson.M, error) {
	var updatedElection bson.M
	idHex, err := primitive.ObjectIDFromHex(id)
	var adminData []*model.User

	data, err := UserCollection.Find(Userctx, bson.D{primitive.E{Key: "email", Value: candidatesRequest.Email}})
	adminData, err = convertDbResultIntoUserStruct(data)
	// adminData, err = convertDbResultIntoUserStruct(data)
	if err != nil {
		return updatedElection, err
	}
	filter := bson.D{}
	flag := true

	if id != "" {
		idHex, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return updatedElection, err
		}
		filter = append(filter, primitive.E{Key: "_id", Value: idHex})
		flag = false
	}
	if flag {
		if candidatesRequest.Email != "" {
			filter = append(filter, primitive.E{Key: "email", Value: bson.M{"$regex": candidatesRequest.Email}})
			flag = false
		}
	}

	var Elections []*model.Election

	cur, err := ElectionCollection.Find(Electionctx, bson.D{primitive.E{Key: "_id", Value: idHex}})

	if err != nil {
		return updatedElection, errors.New("unable to query db")
	}

	for cur.Next(Electionctx) {
		var e model.Election

		err := cur.Decode(&e)

		if err != nil {
			return updatedElection, err
		}

		Elections = append(Elections, &e)
	}
	var Candidates []model.Candidates
	Candidates = Elections[0].Candidates
	// Candidates[].find
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

	// updatedDocument := &model.User{}
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
