package source

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"text/template"
)

// PProfile represents a product's profile
type PProfile struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	ImageUrls   []string `json:"imageUrls"`
	Price       float64  `json:"price"`
	Category    string   `json:"category"`
	PublishDate string   `json:"publishDate"`
	IsNew       bool     `json:"isNew"`
}

// Category represents a product category
type Category struct {
	Title    string
	ImageURL string
}

// CategoryProduct associates products with a category
type CategoryProduct struct {
	MatchedProducts []PProfile
	MatchedCategory Category
}

var (
	Products   []PProfile
	Categories []Category
	NewProduct PProfile
	mutex      sync.Mutex // Protect shared resources
)

func main() {
	// Define the routes
	http.HandleFunc("/", HomePageHandler)
	http.HandleFunc("/admin", AdminPageHandler)
	http.HandleFunc("/admin/dashboard", AdminDashboardPageHandler)
	http.HandleFunc("/admin/dashboard/add/", AddingHandler)
	http.HandleFunc("/admin/dashboard/add/product", AddNewProductPageHandler)
	http.HandleFunc("/menu", MenuPageHandler)
	http.HandleFunc("/about-us", AboutUsPageHandler)
	http.HandleFunc("/product", ProductPageHandler)
	http.HandleFunc("/contact-us", ContactUsPageHandler)
	http.HandleFunc("/geo-localisation", GeoLocalisationPageHandler)

	// Start the server
	http.ListenAndServe(":8080", nil)
}

// renderTemplate is a utility function to execute a template
func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

// HomePageHandler handles requests for the homepage
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}

	Categories = []Category{
		{Title: "Ice Cream", ImageURL: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTjZ5QpVr0Yq0taYNdA7vg1AfD7KUfEbk_NXQ&s"},
		{Title: "Frozen Yogurt", ImageURL: "https://images.immediate.co.uk/production/volatile/sites/30/2020/08/recipe-image-legacy-id-1029452_10-563fda8.jpg"},
		{Title: "Sorbet", ImageURL: "https://www.lecremedelacrumb.com/wp-content/uploads/2014/06/raspberry-sorbet-1.jpg"},
		{Title: "Gelato", ImageURL: "https://emmaduckworthbakes.com/wp-content/uploads/2023/06/Chocolate-Gelato-Recipe.jpg"},
		{Title: "Sundaes", ImageURL: "https://www.keep-calm-and-eat-ice-cream.com/wp-content/uploads/2022/08/Ice-cream-sundae-hero-11.jpg"},
	}

	renderTemplate(w, "./static/home.html", &Categories)
}

// AdminPageHandler handles requests for the admin page
func AdminPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/admin" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}

	renderTemplate(w, "./static/admin.html", nil)
}

// AdminDashboardPageHandler handles requests for the admin dashboard
func AdminDashboardPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/admin/dashboard" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}

	renderTemplate(w, "./static/dashboard.html", nil)
}

// AddNewProductPageHandler handles product addition requests
func AddNewProductPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/admin/dashboard/add/product" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}

	NewProduct = PProfile{} // Reset NewProduct
	renderTemplate(w, "./static/dashboard-add-product.html", NewProduct.ImageUrls)
	// http.re
}

// AddingHandler handles the new product upload
func AddingHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/admin/dashboard/add/" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}

	if r.Method == "POST" {
		if err := r.ParseMultipartForm(10 << 20); err != nil {
			http.Error(w, "file too large", http.StatusInternalServerError)
			return
		}

		// Parse product data
		mutex.Lock()
		NewProduct.ID = len(Products)
		mutex.Unlock()
		NewProduct.Title = r.FormValue("productName")
		NewProduct.Description = r.FormValue("description")
		price, err := strconv.ParseFloat(r.FormValue("price"), 64)
		if err != nil {
			http.Error(w, "invalid price value", http.StatusBadRequest)
			return
		}
		NewProduct.Price = price
		NewProduct.Category = r.FormValue("category")
		NewProduct.PublishDate = r.FormValue("publishDate")
		NewProduct.IsNew = r.FormValue("isNew") == "true"

		// Handle file upload
		files := r.MultipartForm.File["images"]
		if len(files) == 0 {
			http.Error(w, "no files uploaded", http.StatusInternalServerError)
			return
		}

		uploadDir := "./static/uploads/"
		if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
			http.Error(w, "unable to create upload directory", http.StatusInternalServerError)
			return
		}

		var uploadedImages []string
		for _, fileHeader := range files {
			if err := SaveUploadedFile(fileHeader, uploadDir); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			uploadedImages = append(uploadedImages, filepath.Join(uploadDir, fileHeader.Filename))
		}
		NewProduct.ImageUrls = uploadedImages

		// Save product data (with concurrency protection)
		mutex.Lock()
		Products = append(Products, NewProduct)
		mutex.Unlock()

		// Save product details to file
		SaveProductsToFile()

		NewProduct = PProfile{} // Reset for next product
		http.Redirect(w, r, "/admin/dashboard/add/product", http.StatusSeeOther)
	}
}

// MenuPageHandler handles requests for the menu page
func MenuPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/menu" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}

	parsedCategories := createCategoryProducts()

	renderTemplate(w, "./static/menu.html", &parsedCategories)
}

// createCategoryProducts organizes products into categories
func createCategoryProducts() []CategoryProduct {
	parsedCategories := make([]CategoryProduct, len(Categories))
	for i := range Categories {
		parsedCategories[i].MatchedCategory = Categories[i]
		for _, product := range Products {
			if product.Category == Categories[i].Title {
				parsedCategories[i].MatchedProducts = append(parsedCategories[i].MatchedProducts, product)
			}
		}
	}
	return parsedCategories
}

// AboutUsPageHandler handles requests for the "About Us" page
func AboutUsPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about-us" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}

	// You can add specific data here for the "About Us" page if needed
	renderTemplate(w, "./static/about-us.html", nil)
}

// ProductPageHandler handles requests for individual product pages
func ProductPageHandler(w http.ResponseWriter, r *http.Request) {
	productIDStr := r.URL.Query().Get("id")
	if productIDStr == "" {
		http.Error(w, "Product ID not specified", http.StatusBadRequest)
		return
	}

	// Convert product ID to integer
	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	// Find the product by ID
	var selectedProduct *PProfile
	for _, product := range Products {
		if product.ID == productID {
			selectedProduct = &product
			break
		}
	}

	if selectedProduct == nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	// Render the product page template with the selected product data
	renderTemplate(w, "./static/product.html", selectedProduct)
}

// ContactUsPageHandler handles requests for the "Contact Us" page
func ContactUsPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/contact-us" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}

	// Render the contact us page template
	renderTemplate(w, "./static/contact-us.html", nil)
}

// GeoLocalisationPageHandler handles requests for the "Geo Localisation" page
func GeoLocalisationPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/geo-localisation" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}

	// You can add specific data for the Geo Localisation page here if needed
	renderTemplate(w, "./static/geo-localisation.html", nil)
}
