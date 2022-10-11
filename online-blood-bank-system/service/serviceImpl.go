package service

import (
	"bloodBank/model"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/unidoc/unipdf/v3/common/license"
	"github.com/unidoc/unipdf/v3/creator"
	pdfModel "github.com/unidoc/unipdf/v3/model"
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

const dir = "data/download/"

var fileName string

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

	err = license.SetMeteredKey("db722e5d9fa7cb0335b8fd3302096d8c835e8d8f10f4d3d7f6e2b09fb85229e1")
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

func (e *Connection) SaveUserDetails(reqBody model.User) (*mongo.InsertOneResult, error) {

	data, err := CollectionUser.InsertOne(ctx, reqBody)
	if err != nil {
		log.Println(err)
		return data, errors.New("Unable to store data")
	}
	fmt.Println(data)
	return data, nil
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

func (e *Connection) SaveDonorData(donorData model.Donor) (model.Donor, error) {
	// file := "donorData" + fmt.Sprintf("%v", time.Now().Format("3_4_5_pm"))
	donorData.DonationTime = time.Now()
	data, err := CollectionDonor.InsertOne(ctx, donorData)
	if err != nil {
		log.Println(err)
		return donorData, errors.New("Unable to store data")
	}
	fmt.Println(data)
	str, err := updateAvailableBlood(donorData)
	if err != nil {
		log.Println(err)
		return donorData, err
	}

	fmt.Println(str)
	_, err = CertificatesOfBloodDonated(donorData)
	if err != nil {
		log.Println(err)
		return donorData, err
	}
	return donorData, nil
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

func (e *Connection) ApplyBloodPatientDetails(reqBody model.Patient) (model.Patient, error) {

	deduct, err := deductOrAddBloodUnitsFromBloodDetails(reqBody.RequestedUnits, reqBody.Location, "Deduct")
	if err == nil {
		reqBody.IsBloodProvided = true
		// return reqBody, err
	}
	fmt.Println(deduct)

	reqBody.RequestedTime = time.Now()
	data, err := CollectionPatient.InsertOne(ctx, reqBody)
	if err != nil {
		log.Println(err)
		return reqBody, errors.New("Unable to store data")
	}
	CertificatesOfBloodRecieved(reqBody)
	fmt.Println(data)
	return reqBody, nil
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

func CertificatesOfBloodDonated(donorDetails model.Donor) (string, error) {
	file := "BloodDonatationCertificate" + donorDetails.Name + fmt.Sprintf("%v", time.Now().Format("2006-01-02_3_4_5_pm"))
	c := creator.New()
	c.SetPageMargins(20, 20, 20, 20)

	font, err := pdfModel.NewStandard14Font(pdfModel.HelveticaName)
	if err != nil {
		return "", err
	}

	fontBold, err := pdfModel.NewStandard14Font(pdfModel.HelveticaBoldName)
	if err != nil {
		return "", err
	}

	// Generate basic usage chapter.
	if err := basicUsage(c, font, fontBold, donorDetails); err != nil {
		return "", err
	}
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "", err
	}
	err = c.WriteToFile(dir + file + ".pdf")
	if err != nil {
		return "", err
	}
	return "Certificate Download Successfully : " + dir + file + ".pdf", nil
}

func basicUsage(c *creator.Creator, font, fontBold *pdfModel.PdfFont, donorDetails model.Donor) error {
	// Create chapter.
	ch := c.NewChapter("Blood Donatation Certificate")
	ch.SetMargins(0, 0, 10, 0)
	ch.GetHeading().SetFont(font)
	ch.GetHeading().SetFontSize(20)
	ch.GetHeading().SetColor(creator.ColorRGBFrom8bit(72, 86, 95))

	contentAlignH(c, ch, font, fontBold, donorDetails)

	// Draw chapter.
	if err := c.Draw(ch); err != nil {
		return err
	}
	return nil
}

func contentAlignH(c *creator.Creator, ch *creator.Chapter, font, fontBold *pdfModel.PdfFont, donorDetails model.Donor) {

	fontColor := creator.ColorRGBFrom8bit(0, 0, 0)
	normalFontSize := 10.0
	d := c.NewParagraph("Donation Date : " + donorDetails.DonationTime.Format("2006-01-02 3-4-5 pm"))
	d.SetFont(font)
	d.SetFontSize(normalFontSize)
	d.SetColor(fontColor)
	d.SetMargins(0, 0, 10, 0)
	ch.Add(d)
	x := c.NewParagraph("YOU ARE AWESOME!")
	x.SetFont(fontBold)
	x.SetFontSize(14)
	x.SetColor(creator.ColorBlack)
	x.SetMargins(200, 0, 10, 10)
	ch.Add(x)
	z := c.NewParagraph("We are pleased to appriciate the nobel gesture of Mr./Mrs. " + donorDetails.Name + " for his/her voluntary contribution in blood donation.")
	z.SetFont(c.NewTextStyle().Font)
	z.SetFontSize(normalFontSize)
	z.SetColor(creator.ColorBlack)
	z.SetMargins(0, 0, 10, 0)
	ch.Add(z)
	y := c.NewParagraph("Age : " + fmt.Sprintf("%v", donorDetails.Age))
	y.SetFont(font)
	y.SetFontSize(normalFontSize)
	y.SetColor(creator.ColorBlack)
	y.SetMargins(0, 0, 10, 0)
	ch.Add(y)
	b := c.NewParagraph("Blood Group : " + donorDetails.BloodGroup)
	b.SetFont(font)
	b.SetFontSize(normalFontSize)
	b.SetColor(creator.ColorBlack)
	b.SetMargins(0, 0, 10, 0)
	ch.Add(b)
	a := c.NewParagraph("Units : " + fmt.Sprintf("%v", donorDetails.Units))
	a.SetFont(font)
	a.SetFontSize(normalFontSize)
	a.SetColor(creator.ColorBlack)
	a.SetMargins(0, 0, 10, 0)
	ch.Add(a)

	e := c.NewParagraph("Location : " + donorDetails.Location)
	e.SetFont(font)
	e.SetFontSize(normalFontSize)
	e.SetColor(creator.ColorBlack)
	e.SetMargins(0, 0, 10, 0)
	ch.Add(e)
	f := c.NewParagraph("AdharCard : " + donorDetails.Adharcard)
	f.SetFont(font)
	f.SetFontSize(normalFontSize)
	f.SetColor(creator.ColorBlack)
	f.SetMargins(0, 0, 10, 0)
	ch.Add(f)
	m := c.NewParagraph("Authorized Signature : " + "____________")
	m.SetFont(font)
	m.SetFontSize(normalFontSize)
	m.SetColor(creator.ColorBlack)
	m.SetMargins(400, 0, 20, 0)
	ch.Add(m)

}

func CertificatesOfBloodRecieved(patientDetails model.Patient) (string, error) {
	file := "BloodRecievedCertificate" + patientDetails.Name + fmt.Sprintf("%v", time.Now().Format("2006-01-02_3_4_5_pm"))
	c := creator.New()
	c.SetPageMargins(20, 20, 20, 20)

	font, err := pdfModel.NewStandard14Font(pdfModel.HelveticaName)
	if err != nil {
		return "", err
	}

	fontBold, err := pdfModel.NewStandard14Font(pdfModel.HelveticaBoldName)
	if err != nil {
		return "", err
	}

	// Generate basic usage chapter.
	ch := c.NewChapter("Blood Recived Reciept")
	ch.SetMargins(0, 0, 10, 0)
	ch.GetHeading().SetFont(font)
	ch.GetHeading().SetFontSize(20)
	ch.GetHeading().SetColor(creator.ColorRGBFrom8bit(72, 86, 95))

	contentAlignHBloodRecieved(c, ch, font, fontBold, patientDetails)

	// Draw chapter.
	if err := c.Draw(ch); err != nil {
		return "", err
	}

	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "", err
	}
	err = c.WriteToFile(dir + file + ".pdf")
	if err != nil {
		return "", err
	}
	return "Reciept Download Successfully : " + dir + file + ".pdf", nil
}

func contentAlignHBloodRecieved(c *creator.Creator, ch *creator.Chapter, font, fontBold *pdfModel.PdfFont, patientDetails model.Patient) {
	normalFontSize := 10.0
	// normalFontColorGreen := creator.ColorRGBFrom8bit(4, 79, 3)
	d := c.NewParagraph("Recieved Date : " + patientDetails.RequestedTime.Format("2006-01-02 3-4-5 pm"))
	d.SetFont(font)
	d.SetFontSize(normalFontSize)
	d.SetColor(creator.ColorBlack)
	d.SetMargins(0, 0, 10, 0)
	ch.Add(d)

	x := c.NewParagraph("Name : " + patientDetails.Name)
	x.SetFont(font)
	x.SetFontSize(normalFontSize)
	x.SetColor(creator.ColorBlack)
	x.SetMargins(0, 0, 10, 0)
	ch.Add(x)
	y := c.NewParagraph("Age : " + fmt.Sprintf("%v", patientDetails.Age))
	y.SetFont(font)
	y.SetFontSize(normalFontSize)
	y.SetColor(creator.ColorBlack)
	y.SetMargins(0, 0, 10, 0)
	ch.Add(y)
	b := c.NewParagraph("Blood Group : " + patientDetails.BloodGroup)
	b.SetFont(font)
	b.SetFontSize(normalFontSize)
	b.SetColor(creator.ColorBlack)
	b.SetMargins(0, 0, 10, 0)
	ch.Add(b)
	e := c.NewParagraph("Location : " + patientDetails.Location)
	e.SetFont(font)
	e.SetFontSize(normalFontSize)
	e.SetColor(creator.ColorBlack)
	e.SetMargins(0, 0, 10, 0)
	ch.Add(e)
	f := c.NewParagraph("Adharcard : " + patientDetails.Adharcard)
	f.SetFont(font)
	f.SetFontSize(normalFontSize)
	f.SetColor(creator.ColorBlack)
	f.SetMargins(0, 0, 10, 0)
	ch.Add(f)
	g := c.NewParagraph("Authorized Signature : " + "____________")
	g.SetFont(font)
	g.SetFontSize(normalFontSize)
	g.SetColor(creator.ColorBlack)
	g.SetMargins(0, 0, 10, 0)
	ch.Add(g)
	m := c.NewParagraph("Reciver's Signature : " + "____________")
	m.SetFont(font)
	m.SetFontSize(normalFontSize)
	m.SetColor(creator.ColorBlack)
	m.SetMargins(400, 0, 20, 0)
	ch.Add(m)

}
