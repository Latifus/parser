package utils

import (
	"testing"
)

var name string = "file.pdf"

func TestGetIntermediatePDFFileName(t *testing.T) {

	expected := "ocr-file.pdf"
	actual := GetIntermediatePDFFileName(name)

	if expected != actual {
		t.Fatal("Wrong resul")
	}
}


func TestPrepareNameForTxtFile(t *testing.T) {
	expected := "file.txt"
	actual := PrepareNameForTxtFile(name)

	if expected != actual {
		t.Fatal("Wrong resul")
	}
}