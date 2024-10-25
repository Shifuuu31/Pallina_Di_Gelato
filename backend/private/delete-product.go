package private

import (
	"fmt"
	"net/http"

	"github.com/Shifuuu31/Pallina_Di_Gelato/backend"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/dashboard/delete/product" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}

	backend.RenderTemplate(w, backend.Template, "delete-product.html", backend.Products)
}

func DeletingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deleting Handler Called")

	if r.URL.Path != "/dashboard/delete-product" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}

	productID := r.URL.Query().Get("id")
	if productID == "" {
		http.Error(w, "Product ID not specified", http.StatusBadRequest)
		return
	}

	// Locking to protect access to the Products slice
	backend.Mutex.Lock()
	defer backend.Mutex.Unlock()

	var message string
	productFound := false

	// Find and delete the product by ID
	for i, product := range backend.Products {
		if product.ID == productID {
			// Remove the product by slicing
			backend.Products = append(backend.Products[:i], backend.Products[i+1:]...)
			productFound = true
			message = "Product successfully deleted."
			break
		}
	}

	if !productFound {
		message = "Product not found."
	} else {
		// Save the updated product list to file
		backend.SaveProductsToFile()
	}

	// Pass the message to the template and redirect to the delete product page
	http.Redirect(w, r, fmt.Sprintf("/dashboard/delete/product?msg=%s", message), http.StatusSeeOther)
}
