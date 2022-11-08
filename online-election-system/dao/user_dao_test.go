package dao

import (
	"fmt"
	"online-election-system/helper"
	"testing"
)

func TestUploadUserDocuments(t *testing.T) {
	var uploadPath = "upload/userDocuments"
	filePath := "C:/Users/Vidhi/Downloads/529_3_334359_1666357514_Databricks - Generic.pdf"
	msg, err := helper.UploadFile(filePath, uploadPath)

	if err == nil {
		got := msg
		want := "File Downloaded with size :115272"
		fmt.Println("got:", got)
		fmt.Println("want:", want)
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	if err != nil {
		got := msg
		want := ""
		fmt.Println("got:", got)
		fmt.Println("want:", want)
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

}

func TestConvertDbResultIntoUserStruct(t *testing.T) {
	// var fetchDataCursor *mongo.Cursor
	//  var fetchDataCursor * mongo.Cursor := fetchDataCursor.decode( *go.mongodb.org/mongo-driver/mongo.Cursor {Current: go.mongodb.org/mongo-driver/bson.Raw len: 0, cap: 0, nil, bc: go.mongodb.org/mongo-driver/mongo.batchCursor(*go.mongodb.org/mongo-driver/x/mongo/driver.BatchCursor) *{clientSession: *(*"go.mongodb.org/mongo-driver/x/mongo/driver/session.Client")(0xc000326270), clock: *(*"go.mongodb.org/mongo-driver/x/mongo/driver/session.ClusterClock")(0xc0000ae240), comment: (*"go.mongodb.org/mongo-driver/x/bsonx/bsoncore.Value")(0xc0000da010), database: "OnlineElection", collection: "User", id: 0, err: error nil, server: go.mongodb.org/mongo-driver/x/mongo/driver.Server(*go.mongodb.org/mongo-driver/x/mongo/driver/topology.SelectedServer) ..., serverDescription: (*"go.mongodb.org/mongo-driver/mongo/description.Server")(0xc0000da078), errorProcessor: go.mongodb.org/mongo-driver/x/mongo/driver.ErrorProcessor nil, connection: go.mongodb.org/mongo-driver/x/mongo/driver.PinnedConnection nil, batchSize: 0, maxTimeMS: 0, currentBatch: *(*"go.mongodb.org/mongo-driver/x/bsonx/bsoncore.DocumentSequence")(0xc0000aa9c0), firstBatch: true, cmdMonitor: *go.mongodb.org/mongo-driver/event.CommandMonitor nil, postBatchResumeToken: go.mongodb.org/mongo-driver/x/bsonx/bsoncore.Document len: 0, cap: 0, nil, crypt: go.mongodb.org/mongo-driver/x/mongo/driver.Crypt nil, serverAPI: *go.mongodb.org/mongo-driver/x/mongo/driver.ServerAPIOptions nil, limit: 0, numReturned: 1}, batch: *go.mongodb.org/mongo-driver/x/bsonx/bsoncore.DocumentSequence nil, batchLength: 1, registry: *go.mongodb.org/mongo-driver/bson/bsoncodec.Registry {typeEncoders: map[reflect.Type]go.mongodb.org/mongo-driver/bson/bsoncodec.ValueEncoder [...], typeDecoders: map[reflect.Type]go.mongodb.org/mongo-driver/bson/bsoncodec.ValueDecoder [...], interfaceEncoders: []go.mongodb.org/mongo-driver/bson/bsoncodec.interfaceValueEncoder len: 3, cap: 3, [(*"go.mongodb.org/mongo-driver/bson/bsoncodec.interfaceValueEncoder")(0xc00006a180),(*"go.mongodb.org/mongo-driver/bson/bsoncodec.interfaceValueEncoder")(0xc00006a1a0),(*"go.mongodb.org/mongo-driver/bson/bsoncodec.interfaceValueEncoder")(0xc00006a1c0)], interfaceDecoders: []go.mongodb.org/mongo-driver/bson/bsoncodec.interfaceValueDecoder len: 2, cap: 2, [(*"go.mongodb.org/mongo-driver/bson/bsoncodec.interfaceValueDecoder")(0xc000036480),(*"go.mongodb.org/mongo-driver/bson/bsoncodec.interfaceValueDecoder")(0xc0000364a0)], kindEncoders: map[reflect.Kind]go.mongodb.org/mongo-driver/bson/bsoncodec.ValueEncoder [...], kindDecoders: map[reflect.Kind]go.mongodb.org/mongo-driver/bson/bsoncodec.ValueDecoder [...], typeMap: map[go.mongodb.org/mongo-driver/bson/bsontype.Type]reflect.Type [...], mu: (*sync.RWMutex)(0xc00014cc28)}, clientSession: *go.mongodb.org/mongo-driver/x/mongo/driver/session.Client {Server: *(*"go.mongodb.org/mongo-driver/x/mongo/driver/session.Server")(0xc00020a580), ClientID: go.mongodb.org/mongo-driver/internal/uuid.UUID [235,117,52,245,205,72,66,190,191,176,212,40,157,9,216,218], ClusterTime: go.mongodb.org/mongo-driver/bson.Raw len: 0, cap: 0, nil, Consistent: true, OperationTime: *go.mongodb.org/mongo-driver/bson/primitive.Timestamp nil, SessionType: Implicit (1), Terminated: true, RetryingCommit: false, Committing: false, Aborting: false, RetryWrite: false, RetryRead: false, Snapshot: false, CurrentRc: *go.mongodb.org/mongo-driver/mongo/readconcern.ReadConcern nil, CurrentRp: *go.mongodb.org/mongo-driver/mongo/readpref.ReadPref nil, CurrentWc: *go.mongodb.org/mongo-driver/mongo/writeconcern.WriteConcern nil, CurrentMct: *time.Duration nil, transactionRc: *go.mongodb.org/mongo-driver/mongo/readconcern.ReadConcern nil, transactionRp: *go.mongodb.org/mongo-driver/mongo/readpref.ReadPref nil, transactionWc: *go.mongodb.org/mongo-driver/mongo/writeconcern.WriteConcern nil, transactionMaxCommitTime: *time.Duration nil, pool: *(*"go.mongodb.org/mongo-driver/x/mongo/driver/session.Pool")(0xc0000aa420), TransactionState: None (0), PinnedServer: *go.mongodb.org/mongo-driver/mongo/description.Server nil, RecoveryToken: go.mongodb.org/mongo-driver/bson.Raw len: 0, cap: 0, nil, PinnedConnection: go.mongodb.org/mongo-driver/x/mongo/driver/session.LoadBalancedTransactionConnection nil, SnapshotTime: *go.mongodb.org/mongo-driver/bson/primitive.Timestamp nil}, err: error nil})
	// msg, err := convertDbResultIntoUserStruct(fetchDataCursor)

	// 	if err == nil {
	// 		got := msg
	// 		want := "File Downloaded with size :115272"
	// 		fmt.Println("got:", got)
	// 		fmt.Println("want:", want)
	// 		if got != want {
	// 			t.Errorf("got %q want %q", got, want)
	// 		}
	// 	}

	// 	if err != nil {
	// 		got := msg
	// 		want := ""
	// 		fmt.Println("got:", got)
	// 		fmt.Println("want:", want)
	// 		if got != want {
	// 			t.Errorf("got %q want %q", got, want)
	// 		}
	// 	}

}
