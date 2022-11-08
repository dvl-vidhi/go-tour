package dao

import (
	"fmt"
	"online-election-system/helper"
	"testing"
)

func TestUploadVoteSign(t *testing.T) {
	var uploadPath = "upload/voteSign/"
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
