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

type UserDAO struct {
	Server     string
	Database   string
	Collection string
}

var UserCollection *mongo.Collection
var Userctx = context.TODO()

func (e *UserDAO) UserConnect() {
	clientOptions := options.Client().ApplyURI(e.Server)
	client, err := mongo.Connect(Userctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(Userctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	UserCollection = client.Database(e.Database).Collection(e.Collection)
}

func (e *UserDAO) UserInsert(User model.User) (model.User, error) {
	bool, err := validateByNameAndDob(User)

	if err != nil {
		return User, err
	}
	if !bool {
		return User, errors.New("User already present")
	}
	var uploadPath = "upload/userDocuments"
	msg, err := helper.UploadFile(User.UploadedDocs.DocumentPath, uploadPath)
	if err != nil {
		log.Println(err)
		return User, errors.New("Unable to upload file")
	}

	fmt.Println("Upload file:", msg)
	User.IsVerified = false
	data, err := UserCollection.InsertOne(Userctx, User)

	if err != nil {
		fmt.Print(err.Error())
		return User, errors.New("unable to create new record")
	}

	if oid, ok := data.InsertedID.(primitive.ObjectID); ok {

		User.ID = oid

	}

	return User, nil
}

func (e *UserDAO) VerifyUser(req model.User, id string) (bson.M, error) {
	var verifiedUser bson.M
	idHex, err := primitive.ObjectIDFromHex(id)
	var adminData []*model.User

	data, err := UserCollection.Find(Userctx, bson.D{primitive.E{Key: "email", Value: req.Email}})
	adminData, err = convertDbResultIntoUserStruct(data)

	if err != nil {
		return verifiedUser, err
	}

	if len(adminData) == 0 {
		return verifiedUser, mongo.ErrNoDocuments
	}

	filter := bson.D{primitive.E{Key: "_id", Value: idHex}}

	UpdateQuery := bson.D{}
	UpdateQuery = append(UpdateQuery, primitive.E{Key: "is_verified", Value: req.IsVerified})
	UpdateQuery = append(UpdateQuery, primitive.E{Key: "verified_by.id", Value: adminData[0].ID})
	UpdateQuery = append(UpdateQuery, primitive.E{Key: "verified_by.name", Value: adminData[0].Name})
	update := bson.D{{Key: "$set", Value: UpdateQuery}}

	err = UserCollection.FindOneAndUpdate(Userctx, filter, update).Decode(&verifiedUser)

	data, err = UserCollection.Find(Userctx, filter)
	if err != nil {
		return verifiedUser, err
	}
	//Send mail method
	return verifiedUser, nil
}

func (epd *UserDAO) Update(user model.User, id string) (bson.M, error) {
	var updatedUser bson.M
	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return updatedUser, err
	}

	filter := bson.D{primitive.E{Key: "_id", Value: idHex}}

	update := bson.D{primitive.E{Key: "$set", Value: user}}

	// e := &model.User{}
	err = UserCollection.FindOneAndUpdate(Userctx, filter, update).Decode(&updatedUser)
	if err != nil {
		return updatedUser, err
	}

	if updatedUser == nil {
		return updatedUser, errors.New("Data not present in db given by Id or it is deactivated")
	}

	return updatedUser, nil
}

func (e *UserDAO) UserFindById(id string) ([]*model.User, error) {
	var Users []*model.User

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Users, err
	}
	data, err := UserCollection.Find(Userctx, bson.D{primitive.E{Key: "_id", Value: idHex}})

	if err != nil {
		return Users, errors.New("unable to query db")
	}

	if err != nil {
		return Users, err
	}
	Users, err = convertDbResultIntoUserStruct(data)

	if err != nil {
		return Users, err
	}

	if len(Users) == 0 {
		return Users, mongo.ErrNoDocuments
	}

	return Users, nil
}

func (e *UserDAO) FilterOnUsersDetails(req model.UserFilter) ([]*model.User, error) {
	var Users []*model.User
	query := bson.D{}

	if req.Name != "" {
		query = append(query, primitive.E{Key: "name", Value: req.Name})
	}
	if req.Role != "" {
		query = append(query, primitive.E{Key: "role", Value: req.Role})
	}
	if req.IsVerified != false {
		query = append(query, primitive.E{Key: "is_verified", Value: req.IsVerified})
	}
	if req.FatherName != "" {
		query = append(query, primitive.E{Key: "personal_info.father_name", Value: req.FatherName})
	}
	if req.City != "" {
		query = append(query, primitive.E{Key: "personal_info.address.city", Value: req.City})
	}
	if req.State != "" {
		query = append(query, primitive.E{Key: "personal_info.address.state", Value: req.State})
	}

	cur, err := UserCollection.Find(Userctx, query)

	if err != nil {
		return Users, errors.New("unable to query db")
	}

	Users, err = convertDbResultIntoUserStruct(cur)

	if err != nil {
		return Users, err
	}

	if len(Users) == 0 {
		return Users, mongo.ErrNoDocuments
	}

	return Users, nil
}

func (e *UserDAO) UserDelete(id string) (string, error) {

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}
	filter := bson.D{primitive.E{Key: "_id", Value: idHex}}

	res, err := UserCollection.DeleteOne(Userctx, filter)
	if err != nil {
		return "User deletion unsuccessfu!", err
	}

	if res.DeletedCount == 0 {
		return "User deletion unsuccessfu!", errors.New("no Users were deleted")
	}

	return "User deletion successfull", err
}

func (e *UserDAO) UserUpdate(user model.User) (bson.M, error) {
	var updatedUser bson.M
	filter := bson.D{primitive.E{Key: "_id", Value: user.ID}}

	update := bson.D{primitive.E{Key: "$set", Value: user}}

	err := UserCollection.FindOneAndUpdate(Userctx, filter, update).Decode(&updatedUser)
	if err != nil {
		return updatedUser, err
	}

	if updatedUser == nil {
		return updatedUser, errors.New("Data not present in db given by Id or it is deactivated")
	}
	return updatedUser, err
}

func validateByNameAndDob(reqbody model.User) (bool, error) {
	dobStr := reqbody.PersonalInfo.DOB
	fmt.Println(dobStr)
	_, err := UserCollection.Find(Userctx, bson.D{{Key: "name", Value: reqbody.Name}, {Key: "dob", Value: dobStr}})
	if err != nil {
		return false, err
	}
	return true, err
}

func (e *UserDAO) FindByEmailAndPassword(email, password string) ([]*model.User, error) {
	var Users []*model.User

	cur, err := UserCollection.Find(Userctx, bson.D{primitive.E{Key: "email", Value: email}, primitive.E{Key: "password", Value: password}})

	if err != nil {
		return Users, errors.New("unable to query db")
	}

	Users, err = convertDbResultIntoUserStruct(cur)

	if err != nil {
		return Users, err
	}

	if len(Users) == 0 {
		return Users, mongo.ErrNoDocuments
	}

	return Users, nil
}

func convertDbResultIntoUserStruct(fetchDataCursor *mongo.Cursor) ([]*model.User, error) {
	var finaldata []*model.User
	for fetchDataCursor.Next(Userctx) {
		var data model.User
		err := fetchDataCursor.Decode(&data)
		if err != nil {
			return finaldata, err
		}
		finaldata = append(finaldata, &data)
	}
	return finaldata, nil
}
