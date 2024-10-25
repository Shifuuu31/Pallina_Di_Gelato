package private

import (
	"net/http"

	"github.com/Shifuuu31/Pallina_Di_Gelato/backend"
)

// AdminPageHandler handles requests for the admin page
func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}

	backend.RenderTemplate(w, backend.Template, "login.html", nil)
}
