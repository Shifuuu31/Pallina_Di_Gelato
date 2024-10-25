package public

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Shifuuu31/Pallina_Di_Gelato/backend"
)

type UserContact struct {
	ID      string
	Name    string
	Email   string
	Message string
	Date    string
	Seen    bool
}

var User UserContact

// ContactUsPageHandler handles requests for the "Contact Us" page
func ContactUsPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/contact-us" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}

	// Render the contact us page template
	backend.RenderTemplate(w, backend.Template, "contact-us.html", nil)
}

func SubmitForm(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/submit-contact-form" {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}

		User.ID = backend.GenerateUniqueProductID()
		User.Name = r.FormValue("name")
		User.Email = r.FormValue("email")
		User.Message = r.FormValue("message")
		User.Date = time.Now().Format("02-01-2006 00:00:00")
		User.Seen = false

	fmt.Println(User.ID)
	fmt.Println(User.Name)
	fmt.Println(User.Email)
	fmt.Println(User.Message)
	fmt.Println(User.Date)
	fmt.Println(User.Seen)
	http.Redirect(w, r, "/contact-us", http.StatusFound)
}
