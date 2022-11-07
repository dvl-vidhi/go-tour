package model

type UploadedDocs struct {
	DocumentType string `bson:"document_type,omitempty" json:"document_type,omitempty"`
	DocumentIdNo string `bson:"document_id_no,omitempty" json:"document_id_no,omitempty"`
	DocumentPath string `bson:"document_path,omitempty" json:"document_path,omitempty"`
}
