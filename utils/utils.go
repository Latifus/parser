package utils

import "strings"

func GetIntermediatePDFFileName(name string) string {
	return "ocr-" + name
}

func PrepareNameForTxtFile(name string) string {
	index := strings.Index(name, ".")
	var trimmed string
	if index != -1 {
		trimmed = name[:index]
	} else {
		trimmed = name
	}
	return trimmed + ".txt"
}