package public

import (
	"net/http"

	"github.com/Shifuuu31/Palline_Di_Gelato/backend"
)

// ProductPageHandler handles requests for individual product pages
func ProductPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/menu/product" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}
	productID := r.URL.Query().Get("id")
	if productID == "" {
		http.Error(w, "Product ID not specified", http.StatusBadRequest)
		return
	}

	// Find the product by ID
	var selectedProduct *backend.PProfile
	for _, product := range backend.Products {
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
	backend.RenderTemplate(w, backend.Template, "product.html", selectedProduct)
}
