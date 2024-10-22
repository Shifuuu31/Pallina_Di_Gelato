package public

import (
	"net/http"

	"github.com/Shifuuu31/Palline_Di_Gelato/backend"
)

// ContactUsPageHandler handles requests for the "Contact Us" page
func ContactUsPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/contact-us" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}

	// Render the contact us page template
	backend.RenderTemplate(w, backend.Template, "contact-us.html", nil)
}
