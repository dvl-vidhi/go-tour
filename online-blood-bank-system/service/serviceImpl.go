package service

import (
	"bloodBank/model"
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Connection struct {
	Server      string
	Database    string
	Collection1 string
	Collection2 string
	Collection3 string
	Collection4 string
}

var CollectionUser *mongo.Collection
var CollectionDonor *mongo.Collection
var CollectionAvailableBlood *mongo.Collection
var CollectionPatient *mongo.Collection
var ctx = context.TODO()

func (e *Connection) Connect() {
	clientOptions := options.Client().ApplyURI(e.Server)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	CollectionUser = client.Database(e.Database).Collection(e.Collection1)
	CollectionDonor = client.Database(e.Database).Collection(e.Collection2)
	CollectionAvailableBlood = client.Database(e.Database).Collection(e.Collection3)
	CollectionPatient = client.Database(e.Database).Collection(e.Collection4)
}
func (e *Connection) AuthenticateUser(password, UserId string) error {
	var err error
	if (UserId != "") && (password != "") {
		_, err = CollectionUser.Find(ctx, bson.D{primitive.E{Key: "user.user_id", Value: UserId}, primitive.E{Key: "user.password", Value: password}})
	}
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (e *Connection) SaveUserDetails(reqBody model.User) (string, error) {

	data, err := CollectionUser.InsertOne(ctx, reqBody)
	if err != nil {
		log.Println(err)
		return "", errors.New("Unable to store data")
	}
	fmt.Println(data)
	return "User Saved Successfully", nil
}

func (e *Connection) SearchUsersDetailsById(idStr string) ([]*model.User, error) {
	var finalData []*model.User

	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return finalData, err
	}
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	data, err := CollectionUser.Find(ctx, filter)
	if err != nil {
		log.Println(err)
		return finalData, err
	}
	finalData, err = convertDbResultIntoUserStruct(data)
	if err != nil {
		log.Println(err)
		return finalData, err
	}
	return finalData, nil
}

func (e *Connection) UpdateUserDetailsById(reqData model.User, idStr string) (bson.M, error) {
	var updatedDocument bson.M
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return updatedDocument, err
	}
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	UpdateQuery := bson.D{}
	if reqData.Name != "" {
		UpdateQuery = append(UpdateQuery, primitive.E{Key: "name", Value: reqData.Name})
	}
	if reqData.BloodGroup != "" {
		UpdateQuery = append(UpdateQuery, primitive.E{Key: "blood_group", Value: reqData.BloodGroup})
	}
	if reqData.Adharcard != "" {
		UpdateQuery = append(UpdateQuery, primitive.E{Key: "adhar_card", Value: reqData.Adharcard})
	}

	update := bson.D{{Key: "$set", Value: UpdateQuery}}

	r := CollectionUser.FindOneAndUpdate(ctx, filter, update).Decode(&updatedDocument)
	if r != nil {
		return updatedDocument, r
	}
	fmt.Println(updatedDocument)
	if updatedDocument == nil {
		return updatedDocument, errors.New("Data not present in db given by Id or it is deactivated")
	}

	return updatedDocument, nil
}

func (e *Connection) DeleteUserDetailsById(idStr string) (string, error) {
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return "", err
	}
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	res, err := CollectionDonor.DeleteOne(ctx, filter)
	if err != nil {
		return "User deletion unsuccessfu!", err
	}

	if res.DeletedCount == 0 {
		return "User deletion unsuccessfu!", errors.New("no category deleted")
	}

	return "User deletion successfull", err
}

func convertDbResultIntoUserStruct(fetchDataCursor *mongo.Cursor) ([]*model.User, error) {
	var finaldata []*model.User
	for fetchDataCursor.Next(ctx) {
		var data model.User
		err := fetchDataCursor.Decode(&data)
		if err != nil {
			return finaldata, err
		}
		finaldata = append(finaldata, &data)
	}
	return finaldata, nil
}

func convertDate(dateStr string) (time.Time, error) {

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		log.Println(err)
		return date, err
	}
	return date, nil
}

func (e *Connection) SaveDonorData(donorData model.Donor) (string, error) {

	data, err := CollectionDonor.InsertOne(ctx, donorData)
	if err != nil {
		log.Println(err)
		return "", errors.New("Unable to store data")
	}
	fmt.Println(data)
	str, err := updateAvailableBlood(donorData)
	if err != nil {
		log.Println(err)
		return "", err
	}
	fmt.Println(str)
	return "Donor Details Saved Successfully", nil
}

func (e *Connection) SearchDonorDetailsById(idStr string) ([]*model.Donor, error) {
	var finalData []*model.Donor

	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return finalData, err
	}

	data, err := CollectionDonor.Find(ctx, bson.D{primitive.E{Key: "_id", Value: id}})
	if err != nil {
		log.Println(err)
		return finalData, err
	}
	finalData, err = convertDbResultIntoDonorStruct(data)
	if err != nil {
		log.Println(err)
		return finalData, err
	}
	return finalData, nil
}

func (e *Connection) UpdateDonorDetailsById(reqData model.Donor, idStr string) (bson.M, error) {
	var updatedDocument bson.M
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return updatedDocument, err
	}

	UpdateQuery := bson.D{}
	if reqData.Name != "" {
		UpdateQuery = append(UpdateQuery, primitive.E{Key: "name", Value: reqData.Name})
	}

	if reqData.BloodGroup != "" {
		UpdateQuery = append(UpdateQuery, primitive.E{Key: "blood_group", Value: reqData.BloodGroup})
	}
	if reqData.Adharcard != "" {
		UpdateQuery = append(UpdateQuery, primitive.E{Key: "adharcard", Value: reqData.Adharcard})
	}
	if reqData.Location != "" {
		UpdateQuery = append(UpdateQuery, primitive.E{Key: "location", Value: reqData.Location})
	}

	if reqData.Age != 0 {
		UpdateQuery = append(UpdateQuery, primitive.E{Key: "age", Value: reqData.Age})
	}

	update := bson.D{{Key: "$set", Value: UpdateQuery}}
	r := CollectionDonor.FindOneAndUpdate(ctx, bson.D{primitive.E{Key: "_id", Value: id}}, update).Decode(&updatedDocument)
	if r != nil {
		return updatedDocument, r
	}

	if updatedDocument == nil {
		return updatedDocument, errors.New("Data not present in db given by Id or it is deactivated")
	}

	return updatedDocument, nil
}

func (e *Connection) DeleteDonorDetailsById(idStr string) (string, error) {
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return "", err
	}
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	res, err := CollectionDonor.DeleteOne(ctx, filter)
	if err != nil {
		return "Donor deletion unsuccessfu!", err
	}

	if res.DeletedCount == 0 {
		return "Donor deletion unsuccessfu!", errors.New("no category deleted")
	}

	return "Donor deletion successfull", err
}

func convertDbResultIntoDonorStruct(fetchDataCursor *mongo.Cursor) ([]*model.Donor, error) {
	var finaldata []*model.Donor
	for fetchDataCursor.Next(ctx) {
		var data model.Donor
		err := fetchDataCursor.Decode(&data)
		if err != nil {
			return finaldata, err
		}
		finaldata = append(finaldata, &data)
	}
	return finaldata, nil
}

func updateAvailableBlood(donorData model.Donor) (string, error) {
	var finalData []*model.AvailableBlood

	filter := bson.D{
		{Key: "$and",
			Value: bson.A{
				bson.D{primitive.E{Key: "location", Value: donorData.Location}},
				bson.D{primitive.E{Key: "blood_group", Value: donorData.BloodGroup}},
			},
		},
	}
	data, err := CollectionAvailableBlood.Find(ctx, filter)

	finalData, err = convertIntoAvailableBlood(data)
	if err != nil {
		return "", nil
	}
	if finalData == nil {
		saved, err := createNewEntryIntoBloodDetails(donorData, donorData.Units, donorData.DonationTime)
		if err != nil {
			return "", err
		}
		fmt.Println(saved)
	} else {
		unitDB := finalData[0].Units
		addUnit := unitDB + donorData.Units
		fmt.Println("Total Units:", addUnit)
		CollectionAvailableBlood.FindOneAndUpdate(ctx, filter, bson.D{{Key: "$set", Value: bson.D{{Key: "units", Value: addUnit}}}})
	}
	return "Blood Details Saved Successfully", nil
}

func createNewEntryIntoBloodDetails(reqBody model.Donor, unitInt int, depositDate time.Time) (string, error) {
	var bloodDetails model.AvailableBlood

	bloodDetails.Units = unitInt
	bloodDetails.Location = reqBody.Location
	bloodDetails.BloodGroup = reqBody.BloodGroup
	_, err := CollectionAvailableBlood.InsertOne(ctx, bloodDetails)
	if err != nil {
		log.Println(err)
		return "", nil
	}
	return "New entry created in blood details", nil
}

func convertIntoAvailableBlood(fetchDataCursor *mongo.Cursor) ([]*model.AvailableBlood, error) {
	var finaldata []*model.AvailableBlood
	for fetchDataCursor.Next(ctx) {
		var data model.AvailableBlood
		err := fetchDataCursor.Decode(&data)
		if err != nil {
			return finaldata, err
		}
		finaldata = append(finaldata, &data)
	}
	return finaldata, nil
}

func convertUnitsStringIntoInt(units string) (int, error) {
	unitReplace := strings.ReplaceAll(units, "ml", "")
	unitInt, err := strconv.Atoi(unitReplace)
	if err != nil {
		fmt.Println(err)
		return 0, nil
	}
	return unitInt, nil
}

func deductOrAddBloodUnitsFromBloodDetails(units int, location, methodCall string) (string, error) {

	filter := bson.D{
		{Key: "$and",
			Value: bson.A{
				bson.D{{Key: "location", Value: location}},
			},
		},
	}
	fmt.Println(filter)
	data, err := CollectionAvailableBlood.Find(ctx, filter)
	finalData, err := convertIntoAvailableBlood(data)
	fmt.Println(finalData)
	if err != nil {
		return "", nil
	}
	if finalData == nil {
		return "", errors.New("Data not present in Blood details according to given location and desposited date")
	}
	if methodCall == "Deduct" {
		unit := finalData[0].Units
		if !(unit >= units) {
			return "", errors.New("Insufficient Blood!")
		}
		addUnit := unit - units
		fmt.Println("Total Units:", addUnit)
		CollectionAvailableBlood.FindOneAndUpdate(ctx, filter, bson.D{{Key: "$set", Value: bson.D{{Key: "units", Value: addUnit}}}})
		return "Blood units Deduct Successfully", nil
	} else if methodCall == "Add" {
		unit := finalData[0].Units
		addUnit := unit + units
		fmt.Println("Total Units:", addUnit)
		CollectionAvailableBlood.FindOneAndUpdate(ctx, filter, bson.D{{Key: "$set", Value: bson.D{{Key: "units", Value: addUnit}}}})
		return "Blood Units Added Successfully", nil
	}
	return "", nil
}

func (e *Connection) ApplyBloodPatientDetails(reqBody model.Patient) (string, error) {

	deduct, err := deductOrAddBloodUnitsFromBloodDetails(reqBody.RequestedUnits, reqBody.Location, "Deduct")
	if err != nil {
		return "", err
	}
	fmt.Println(deduct)
	reqBody.IsBloodProvided = true
	data, err := CollectionPatient.InsertOne(ctx, reqBody)
	if err != nil {
		log.Println(err)
		return "", errors.New("Unable to store data")
	}
	fmt.Println(data)
	return "Patient Saved Successfully", nil
}

func (e *Connection) GivenBloodPatientDetailsById(idStr string) (bson.M, error) {
	var updatedDocument bson.M
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return updatedDocument, err
	}
	filter := bson.D{primitive.E{Key: "_id", Value: id}}

	UpdateQuery := bson.D{}
	UpdateQuery = append(UpdateQuery, primitive.E{Key: "requested_time", Value: time.Now()})

	update := bson.D{{Key: "$set", Value: UpdateQuery}}

	r := CollectionPatient.FindOneAndUpdate(ctx, filter, update).Decode(&updatedDocument)
	if r != nil {
		return updatedDocument, r
	}
	fmt.Println(updatedDocument)
	if updatedDocument == nil {
		return updatedDocument, errors.New("Data not present in db given by Id or it is deactivated")
	}

	return updatedDocument, nil
}

func (e *Connection) SearchFilterBloodDetails(search model.AvailableBlood) ([]*model.AvailableBlood, error) {
	var searchData []*model.AvailableBlood

	filter := bson.D{}

	if search.BloodGroup != "" {
		filter = append(filter, primitive.E{Key: "blood_group", Value: bson.M{"$regex": search.BloodGroup}})
	}
	if search.Location != "" {
		filter = append(filter, primitive.E{Key: "location", Value: bson.M{"$regex": search.Location}})
	}

	result, err := CollectionAvailableBlood.Find(ctx, filter)
	if err != nil {
		return searchData, err
	}
	data, err := convertIntoAvailableBlood(result)
	if err != nil {
		return searchData, err
	}

	return data, nil
}
