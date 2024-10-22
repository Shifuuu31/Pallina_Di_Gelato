package private

import (
	"net/http"

	"github.com/Shifuuu31/Palline_Di_Gelato/backend"
)

func AddCategory(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/dashboard/add/category" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}

	backend.RenderTemplate(w, backend.Template, "add-category.html", nil)
}
