package public

import (
	"net/http"

	"github.com/Shifuuu31/Palline_Di_Gelato/backend"
)

// AboutUsPageHandler handles requests for the "About Us" page
func AboutUsPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about-us" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}

	// You can add specific data here for the "About Us" page if needed
	backend.RenderTemplate(w, backend.Template, "about-us.html", nil)
}
