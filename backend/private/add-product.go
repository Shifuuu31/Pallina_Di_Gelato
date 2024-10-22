package private

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/Shifuuu31/Palline_Di_Gelato/backend"
	"github.com/google/uuid"
)

// AddNewProductPageHandler handles product addition requests
func AddProductPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/dashboard/add/product" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}

	backend.RenderTemplate(w, backend.Template, "add-product.html", backend.NewProduct.ImageUrls)
}

// AddingHandler handles the new product upload
func AddingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			backend.RenderTemplate(w, backend.Template, "add-product.html", "Product upload failed: file too large")
			return
		}

		// Parse product data and generate unique ID
		backend.NewProduct.ID = generateUniqueProductID()
		backend.NewProduct.Title = r.FormValue("productName")
		backend.NewProduct.Description = r.FormValue("description")
		price, err := strconv.ParseFloat(r.FormValue("price"), 64)
		if err != nil {
			backend.RenderTemplate(w, backend.Template, "add-product.html", "Invalid price value")
			return
		}
		backend.NewProduct.Price = price
		backend.NewProduct.Category = r.FormValue("category")
		parsedTime, err := time.Parse("2006-01-02T15:04", r.FormValue("publishDate"))
		if err != nil {
			backend.RenderTemplate(w, backend.Template, "add-product.html", fmt.Sprintf("Error parsing time: %v", err))
			return
		}
		backend.NewProduct.PublishDate = parsedTime.Format("2006-01-02 15:04:05")
		backend.NewProduct.CreationDate = time.Now().Format("2006-01-02 15:04:05")
		backend.NewProduct.IsNew = r.FormValue("isNew") == "true"
		backend.NewProduct.IsVisible = r.FormValue("isVisible") == "true" // Set visibility based on checkbox

		// Handle file upload
		files := r.MultipartForm.File["images"]
		if len(files) == 0 {
			backend.RenderTemplate(w, backend.Template, "add-product.html", "No files uploaded")
			return
		}

		uploadDir := "./assets/uploads/"
		if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
			backend.RenderTemplate(w, backend.Template, "add-product.html", "Unable to create upload directory")
			return
		}

		var uploadedImages []string
		for _, fileHeader := range files {
			if err := backend.SaveUploadedFile(fileHeader, uploadDir); err != nil {
				backend.RenderTemplate(w, backend.Template, "add-product.html", err.Error())
				return
			}
			uploadedImages = append(uploadedImages, filepath.Join(uploadDir, fileHeader.Filename))
		}
		backend.NewProduct.ImageUrls = uploadedImages

		// Save product data with concurrency protection
		backend.Mutex.Lock()
		backend.Products = append(backend.Products, backend.NewProduct)
		backend.Mutex.Unlock()

		// Save product details to file
		if err := backend.SaveProductsToFile(); err != nil {
			backend.RenderTemplate(w, backend.Template, "add-product.html", "Failed to save product to file")
			return
		}

		backend.NewProduct = backend.PProfile{} // Reset for next product
		backend.RenderTemplate(w, backend.Template, "add-product.html", "Product added successfully!")
		return
	}
	backend.NewProduct = backend.PProfile{} // Reset NewProduct

	http.Redirect(w, r, "/dashboard/add/product", http.StatusSeeOther)
}

func generateUniqueProductID() string {
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
	backend.Mutex.Lock()
	defer backend.Mutex.Unlock()

	for _, product := range backend.Products {
		if product.ID == id {
			return true
		}
	}
	return false
}
