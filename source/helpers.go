package source

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

func LoadProduct() error {
	const filePath string = "/root/Desktop/Palline_Di_Gelato/product_details.json"
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(strings.NewReader(string(jsonData)))
	if err := decoder.Decode(&Products); err != nil {
		// http.Error(w, "Error product details cannot be displayed", http.StatusInternalServerError)
		return err
	}
	return nil
}

// saveProductsToFile saves the product data to a JSON file
func SaveProductsToFile() {
	const filePath string = "/root/Desktop/Palline_Di_Gelato/product_details.json"
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(Products); err != nil {
		fmt.Println("Error encoding products to file:", err)
		return
	}

	// Post-processing to fix escape sequences
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	modifiedContent := strings.ReplaceAll(string(content), "\\u0026", "&")
	if err := os.WriteFile(filePath, []byte(modifiedContent), 0o644); err != nil {
		fmt.Println("Error writing modified content:", err)
	}
}

// saveUploadedFile saves the uploaded file to the specified directory
func SaveUploadedFile(fileHeader *multipart.FileHeader, uploadDir string) error {
	file, err := fileHeader.Open()
	if err != nil {
		return fmt.Errorf("error retrieving the file: %v", err)
	}
	defer file.Close()

	filePath := filepath.Join(uploadDir, fileHeader.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("unable to save the file: %v", err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return fmt.Errorf("failed to save the file: %v", err)
	}
	return nil
}
