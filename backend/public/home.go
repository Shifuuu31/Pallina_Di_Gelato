package public

import (
	"net/http"

	"github.com/Shifuuu31/Pallina_Di_Gelato/backend"
)

// HomePageHandler handles requests for the homepage
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}

	backend.RenderTemplate(w, backend.Template, "home.html", &backend.Categories)
}
