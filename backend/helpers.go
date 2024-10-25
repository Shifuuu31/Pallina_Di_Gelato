package backend

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/google/uuid"
)

// RenderTemplate renders a template and handles any errors.
func RenderTemplate(w http.ResponseWriter, tmpl *template.Template, filename string, data interface{}) {
	w.Header().Set("Content-Type", "text/html")

	err := tmpl.ExecuteTemplate(w, filename, data)
	if err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
		return // Make sure to return to prevent further writes
	}
}

func LoadProducts() error {
	const filePath string = "./database/products.json"
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

func LoadCategories() error {
	const filePath string = "./database/categories.json"
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(strings.NewReader(string(jsonData)))
	if err := decoder.Decode(&Categories); err != nil {
		// http.Error(w, "Error product details cannot be displayed", http.StatusInternalServerError)
		return err
	}
	// fmt.Println(Categories)
	return nil
}

// saveProductsToFile saves the product data to a JSON file
func SaveProductsToFile() error {
	const filePath string = "./database/products.json"
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(Products); err != nil {
		return err
	}

	// Post-processing to fix escape sequences
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	modifiedContent := strings.ReplaceAll(string(content), "\\u0026", "&")
	if err := os.WriteFile(filePath, []byte(modifiedContent), 0o644); err != nil {
		return err
	}
	return nil
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

func GenerateUniqueProductID() string {
	var newID string
	exists := true

	for exists {
		newID = uuid.New().String()
		exists = checkDuplicateID(newID)
	}
	return newID
}

// Checks if the generated ID already exists in the product list
func checkDuplicateID(id string) bool {
	Mutex.Lock()
	defer Mutex.Unlock()

	for _, product := range Products {
		if product.ID == id {
			return true
		}
	}
	return false
}
