package main

import (
	"fmt"
	"io/ioutil"
	"mypdfprojectf/internal"
	"mypdfprojectf/utils"
	"os"
	"os/exec"
)

func main() {
	inputPDF := "w-3.pdf"
	ConnectPdf(inputPDF)
}

func ConnectPdf(inputPDF string) (string, error) {
	intermediatePDF := utils.GetIntermediatePDFFileName(inputPDF)
	outputText := utils.PrepareNameForTxtFile(inputPDF)

	err := applyOCRToPDF(inputPDF, intermediatePDF)
	if err != nil {
		internal.Logger().Error("error handler", "error", err)
		return "", fmt.Errorf("error applying OCR to PDF: %v", err)
	}

	text, err := extractTextFromPDF(intermediatePDF, outputText)
	if err != nil {
		internal.Logger().Error("error handler", "error", err)
		return "", fmt.Errorf("error extracting text from PDF: %v", err)
	}

	fmt.Println("Text extracted successfully!")
	return text, nil
}

func applyOCRToPDF(inputPDF, outputPDF string) error {
	cmd := exec.Command("ocrmypdf", "--force-ocr", inputPDF, outputPDF)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error executing OCR command for input file '%s' and output file '%s': %v", inputPDF, outputPDF, err)
	}
	return nil
}

func extractTextFromPDF(inputPDF, outputTextPath string) (string, error) {
	cmd := exec.Command("pdftotext", inputPDF, outputTextPath)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("error executing pdftotext command for input file '%s' and output file '%s': %v", inputPDF, outputTextPath, err)
	}

	text, err := ioutil.ReadFile(outputTextPath)
	if err != nil {
		return "", fmt.Errorf("error was received while reading the file: "+outputTextPath+" %v\n", err)
	}

	return string(text), nil
}
