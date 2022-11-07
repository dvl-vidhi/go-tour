package dao

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"
	"strings"

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

const uploadPath = "upload/"

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
	// var data []*model.User
	bool, err := validateByNameAndDob(User)

	if err != nil {
		return User, err
	}
	if !bool {
		return User, errors.New("User already present")
	}

	msg, err := UploadFile(User.UploadedDocs.DocumentPath)
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

	// data, err := CollectionUser.InsertOne(ctx, reqBody)
	// if err != nil {
	// 	log.Println(err)
	// 	return reqBody, errors.New("Unable to store data")
	// }
	if oid, ok := data.InsertedID.(primitive.ObjectID); ok {

		User.ID = oid

	}
	// fmt.Println(data)
	// return reqBody, nil

	return User, nil
}

func (e *UserDAO) VerifyUser(req model.User, id string) (bson.M, error) {
	var verifiedUser bson.M
	var adminData []*model.User

	data, err := UserCollection.Find(Userctx, bson.D{primitive.E{Key: "email", Value: req.Email}})
	adminData, err = convertDbResultIntoUserStruct(data)
	// adminData, err = convertDbResultIntoUserStruct(data)
	if err != nil {
		return verifiedUser, err
	}
	filter := bson.D{}
	flag := true

	if id != "" {
		idHex, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return verifiedUser, err
		}
		filter = append(filter, primitive.E{Key: "_id", Value: idHex})
		flag = false
	}
	if flag {
		if req.Email != "" {
			filter = append(filter, primitive.E{Key: "email", Value: bson.M{"$regex": req.Email}})
			flag = false
		}
	}
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
		log.Println(err)
		return Users, err
	}
	Users, err = convertDbResultIntoUserStruct(data)
	if err != nil {
		log.Println(err)
		return Users, err
	}
	return Users, nil
}

func (e *UserDAO) FilterOnUsersDetails(req model.User) ([]*model.User, error) {
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
	if req.PhoneNumber != "" {
		query = append(query, primitive.E{Key: "phone_number", Value: req.PhoneNumber})
	}
	if req.PersonalInfo.FatherName != "" {
		query = append(query, primitive.E{Key: "personal_info.father_name", Value: req.PersonalInfo.FatherName})
	}
	if req.PersonalInfo.Address.City != "" {
		query = append(query, primitive.E{Key: "personal_info.address.city", Value: req.PersonalInfo.Address.City})
	}
	if req.PersonalInfo.Address.Street != "" {
		query = append(query, primitive.E{Key: "personal_info.address.street", Value: req.PersonalInfo.Address.Street})
	}
	if req.PersonalInfo.Address.State != "" {
		query = append(query, primitive.E{Key: "personal_info.address.state", Value: req.PersonalInfo.Address.State})
	}
	if req.PersonalInfo.Address.Country != "" {
		query = append(query, primitive.E{Key: "personal_info.address.country", Value: req.PersonalInfo.Address.Country})
	}

	cur, err := UserCollection.Find(Userctx, query)

	if err != nil {
		return Users, errors.New("unable to query db")
	}

	for cur.Next(Userctx) {
		var e model.User

		err := cur.Decode(&e)

		if err != nil {
			return Users, err
		}

		Users = append(Users, &e)
	}

	if err := cur.Err(); err != nil {
		return Users, err
	}

	cur.Close(Userctx)

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

func (epd *UserDAO) UserUpdate(user model.User) (bson.M, error) {
	var updatedUser bson.M
	filter := bson.D{primitive.E{Key: "_id", Value: user.ID}}

	update := bson.D{primitive.E{Key: "$set", Value: user}}

	// updatedDocument := &model.User{}
	err := UserCollection.FindOneAndUpdate(Userctx, filter, update).Decode(&updatedUser)
	if err != nil {
		return updatedUser, err
	}
	fmt.Println(updatedUser)
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

func UploadFile(path string) (string, error) {
	err := os.MkdirAll(uploadPath, os.ModePerm)
	if err != nil {
		return "", err
	}

	fileURL, err := url.Parse(path)
	if err != nil {
		return "", err
	}
	segments := strings.Split(fileURL.Path, "/")
	fileName := segments[len(segments)-1]
	fileName = uploadPath + fileName
	// Create blank file
	file, err := os.Create(fileName)
	if err != nil {
		return "", err
	}

	resp, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer resp.Close()
	size, err := io.Copy(file, resp)
	defer file.Close()
	return "File Downloaded with size :" + fmt.Sprintf("%v", size), nil
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

func (e *UserDAO) FindByEmailAndPassword(email, password string) ([]*model.User, error) {
	var Users []*model.User

	cur, err := UserCollection.Find(Userctx, bson.D{primitive.E{Key: "email", Value: email}, primitive.E{Key: "password", Value: password}})

	if err != nil {
		return Users, errors.New("unable to query db")
	}

	for cur.Next(Userctx) {
		var e model.User

		err := cur.Decode(&e)

		if err != nil {
			return Users, err
		}

		Users = append(Users, &e)
	}

	if err := cur.Err(); err != nil {
		return Users, err
	}

	cur.Close(Userctx)

	if len(Users) == 0 {
		return Users, mongo.ErrNoDocuments
	}

	return Users, nil
}
