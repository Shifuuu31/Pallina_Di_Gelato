package public

import (
	"net/http"

	"github.com/Shifuuu31/Pallina_Di_Gelato/backend"
)

// FindUsPageHandler handles requests for the "Geo Localisation" page
func FindUsPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/find-us" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}

	// You can add specific data for the Geo Localisation page here if needed
	backend.RenderTemplate(w, backend.Template, "find-us.html", nil)
}
