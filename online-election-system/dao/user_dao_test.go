package dao

import (
	"fmt"
	"testing"
)

func TestUploadFile(t *testing.T) {
	filePath := "C:/Users/Vidhi/Downloads/529_3_334359_1666357514_Databricks - Generic.pdf"
	msg, err := UploadFile(filePath)

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
