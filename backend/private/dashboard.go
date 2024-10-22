package private

import (
	"net/http"

	"github.com/Shifuuu31/Palline_Di_Gelato/backend"
)



// DashboardPageHandler handles requests for the admin dashboard
func DashboardPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/dashboard" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}

	backend.RenderTemplate(w, backend.Template, "dashboard.html", nil)
}
