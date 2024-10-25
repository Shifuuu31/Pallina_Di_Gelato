package private

import (
	"net/http"

	"github.com/Shifuuu31/Pallina_Di_Gelato/backend"
)

func EditProduct(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/dashboard/edit/product" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}

	backend.RenderTemplate(w, backend.Template, "edit-product.html", nil)
}
